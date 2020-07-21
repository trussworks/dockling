import requests

def get_example():
    response = requests.get('http://example.org')
    response.raise_for_status()
    print("Example status_code: ", response.status_code)
