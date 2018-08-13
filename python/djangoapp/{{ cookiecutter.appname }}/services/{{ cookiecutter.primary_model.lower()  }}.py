from ..models import {{  cookiecutter.primary_model }}


def get_{{  cookiecutter.primary_model.lower() }}_by_pk({{  cookiecutter.primary_model.lower() }}_pk: int) -> {{ cookiecutter.primary_model }}:
    """
    Gets a {{  cookiecutter.primary_model.lower() }} by it's pk.

    Args:
      {{  cookiecutter.primary_model.lower() }}_pk: The pk for the instance

    Returns:
      {{  cookiecutter.primary_model }}
    """
    return {{  cookiecutter.primary_model.lower() }}.objects.get(pk={{  cookiecutter.primary_model.lower() }}_pk)