def get_metadata(endpoint):
    return [("caller", "{{ cookiecutter.svc_name }}"), ("endpoint", endpoint)]
