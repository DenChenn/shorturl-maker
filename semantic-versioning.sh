#!/bin/bash

# Change deployment image name
yq -i '.spec.template.spec.containers[0].image = "asia.gcr.io/${{ env.PROJECT_ID }}/${{ needs.build-backend-image.outputs.MAJOR_VERSION}}/${{ env.IMAGE_NAME }}:${{ needs.build-backend-image.outputs.RELEASE_VERSION }}"' ./k8s-config/backend.yaml

#---------------- Deployment ----------------
# change deployment name
yq -i '.metadata.name = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend.yaml

# change deployment label name
yq -i '.metadata.labels.deployment = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend.yaml


# change deployment match pod name
yq -i '.spec.selector.matchLabels.pod = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend.yaml

# change pod label name
yq -i '.spec.template.metadata.labels.pod = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend.yaml

# change pod name
yq -i '.spec.template.metadata.labels.name = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend.yaml

#---------------- Service ----------------
# change service name
yq -i '.metadata.name = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend-service.yaml

# change service match pod name
yq -i '.spec.selector.pod = "url-shortener-backend-${{ needs.build-backend-image.outputs.MAJOR_VERSION}}"' ./k8s-config/backend-service.yaml
