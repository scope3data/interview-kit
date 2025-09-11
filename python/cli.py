import click
import logging
from typing import Optional

from api import Client

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

_client: Optional[Client] = None

def get_client() -> Client:
    """Initialize and return the API client."""
    global _client
    if _client is None:
        _client = Client()
    return _client

@click.group()
def cli():
    get_client()

@cli.command()
def probe():
    """Tests that the API is reachable and API_KEY is set"""
    client = get_client()
    try:
        response = client.measure(["yahoo.com"], "2025-05-01")
        click.echo(str(response))
    except Exception as e:
        click.echo(str(e), err=True)
        exit(1)

@cli.command()
def trends():
    """Fetches the top trending domains"""
    click.echo("Not implemented!")

@cli.command()
@click.argument('properties', nargs=-1, required=True)
@click.option('-d', '--date', help='Date to compare against')
def compare(properties, date):
    """Compares emission data across properties"""
    options = {'date': date}
    result = {'properties': list(properties), 'options': options}
    click.echo(f"{result} compare arguments!")

def execute():
    """Entry point for the CLI"""
    try:
        cli()
    except Exception as e:
        click.echo(str(e), err=True)
        exit(1)

if __name__ == '__main__':
    execute()
