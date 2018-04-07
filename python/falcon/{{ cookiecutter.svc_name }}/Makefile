default: venv docker

upb:
	docker-compose up --build

upd:
	docker-compose up -d

ipyshell:
	{{ cookiecutter.venv_path }}/bin/ipython

venv:
	echo "Creating venv {{ cookiecutter.venv_path }}/{{ cookiecutter.svc_name }}"
	python3 -m venv {{ cookiecutter.venv_path }}
	{{ cookiecutter.venv_path }}/bin/pip install --upgrade pip
	echo "Installing dev dependencies"
	{{ cookiecutter.venv_path }}/bin/pip install -r requirements-dev.txt

docker:
	echo "Building base Docker image (not pushing to remote)"
	./scripts/build_and_push_base_image.sh -v 1.0
