import json

from ..utils import get_logger


logger = get_logger()


def get_data_from_request(request):
    """
    Deserializes the JSON in the request body and returns the result

    Args:
        request (falcon.Request)

    Returns:
        [dict, list]: The deserialized JSON. If nothing can be read returns an empty dict
    """
    data = {}
    try:
        data = json.loads(request.stream.read())
    except ValueError as err:
        logger.warn("error parsing JSON: %s", err)
    return data
