# Copyright (c) 2018 Cisco and/or its affiliates.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This file includes definitions for Docker images used by the Makefile
# and docker build infrastructure. It also contains the targets to build
# and push Docker images

DOCKER_NSE_SSVPP=networkservicemesh/nse-ssvpp
DOCKER_VPP_BASE=networkservicemesh/vpp-base
DOCKER_STRONGSWAN_BASE=networkservicemesh/strongswan-base
DOCKER_SSVPP=networkservicemesh/ssvpp

#
# Targets to build docker images
#
# NOTE: ${COMMIT} is set in .travis.yml from the first 8 bytes of
# ${TRAVIS_COMMIT}. Thus, for travis-ci builds, we tag the Docker images
# with both the name and this first 8 bytes of the commit hash.
#
.PHONY: docker-build-nse-ssvpp
docker-build-nse-ssvpp:
	@${DOCKERBUILD} -t ${DOCKER_NSE_SSVPP} -f build/Dockerfile.nse-ssvpp ../..
	@if [ "x${COMMIT}" != "x" ] ; then \
		docker tag ${DOCKER_NSE_SSVPP} ${DOCKER_NSE_SSVPP}:${COMMIT} ;\
	fi

.PHONY: docker-build-vpp-base
docker-build-vpp-base:
	@${DOCKERBUILD} -t ${DOCKER_VPP_BASE} -f build/Dockerfile.vpp-base ../..
	@if [ "x${COMMIT}" != "x" ] ; then \
		docker tag ${DOCKER_VPP_BASE} ${DOCKER_VPP_BASE}:${COMMIT} ;\
	fi

.PHONY: docker-build-strongswan-base
docker-build-strongswan-base: docker-build-vpp-base
	@${DOCKERBUILD} -t ${DOCKER_STRONGSWAN_BASE} -f build/Dockerfile.strongswan-base ../..
	@if [ "x${COMMIT}" != "x" ] ; then \
		docker tag ${DOCKER_STRONGSWAN_BASE} ${DOCKER_STRONGSWAN_BASE}:${COMMIT} ;\
	fi

.PHONY: docker-build-ssvpp
docker-build-ssvpp: docker-build-strongswan-base
	@${DOCKERBUILD} -t ${DOCKER_SSVPP} -f build/Dockerfile.ssvpp ../..
	@if [ "x${COMMIT}" != "x" ] ; then \
		docker tag ${DOCKER_SSVPP} ${DOCKER_SSVPP}:${COMMIT} ;\
	fi

#
# Targets to push docker images
#
# NOTE: These assume that ${COMMIT} is set and are meant to be called from travis-ci only.
#
.PHONY: docker-login
docker-login:
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin

.PHONY: docker-push-nse-ssvpp
docker-push-nse-ssvpp: docker-login
	docker tag ${DOCKER_NSE_SSVPP}:${COMMIT} ${DOCKER_NSE_SSVPP}:${TAG}
	docker tag ${DOCKER_NSE_SSVPP}:${COMMIT} ${DOCKER_NSE_SSVPP}:travis-${TRAVIS_BUILD_NUMBER}
	docker push ${DOCKER_NSE_SSVPP}

.PHONY: docker-push-vpp-base
docker-push-vpp-base: docker-login
	docker tag ${DOCKER_VPP_BASE}:${COMMIT} ${DOCKER_VPP_BASE}:${TAG}
	docker tag ${DOCKER_VPP_BASE}:${COMMIT} ${DOCKER_VPP_BASE}:travis-${TRAVIS_BUILD_NUMBER}
	docker push ${DOCKER_VPP_BASE}

.PHONY: docker-push-strongswan-base
docker-push-strongswan-base: docker-login
	docker tag ${DOCKER_STRONGSWAN_BASE}:${COMMIT} ${DOCKER_STRONGSWAN_BASE}:${TAG}
	docker tag ${DOCKER_STRONGSWAN_BASE}:${COMMIT} ${DOCKER_STRONGSWAN_BASE}:travis-${TRAVIS_BUILD_NUMBER}
	docker push ${DOCKER_STRONGSWAN_BASE}

.PHONY: docker-push-ssvpp
docker-push-vpp: docker-login
	docker tag ${DOCKER_SSVPP}:${COMMIT} ${DOCKER_SSVPP}:${TAG}
	docker tag ${DOCKER_SSVPP}:${COMMIT} ${DOCKER_SSVPP}:travis-${TRAVIS_BUILD_NUMBER}
	docker push ${DOCKER_SSVPP}
