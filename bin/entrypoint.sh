#!/bin/sh -l

cd $GITHUB_WORKSPACE/$INPUT_ROOT_DIR
tree
terraform init
terraform $INPUT_ACTION