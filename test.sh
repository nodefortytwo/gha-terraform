docker run --rm \
--workdir /github/workspace \
-e GITHUB_WORKSPACE=/github/workspace \
-e INPUT_ROOT_DIR=./terraform \
-e INPUT_BACKEND_CONFIG=remote.config \
-e INPUT_VARS_FILE=vars.tfvars \
-e INPUT_ACTION="validate,plan" \
--mount type=bind,source="$(pwd)"/tests/workspace,target=/github/workspace \
-it $(docker build -q .)