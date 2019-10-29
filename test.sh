docker run --rm \
-e GITHUB_WORKSPACE=/workspace \
-e INPUT_ROOT_DIR=./terraform \
-e INPUT_ACTION=plan \
--mount type=bind,source="$(pwd)"/tests/workspace,target=/workspace \
-it $(docker build -q .)