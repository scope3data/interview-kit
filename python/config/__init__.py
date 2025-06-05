import os
import logging
from typing import Optional
from dotenv import load_dotenv

class Config:
    """Configuration class that loads settings from environment variables and .env file.

    This class follows the singleton pattern to ensure only one instance of the
    configuration exists throughout the application.

    Attributes:
        api_key (str): The Scope3 API key for authentication
        log_level (str): The logging level (default: 'DEBUG')
    """
    _instance: Optional['Config'] = None

    def __new__(cls) -> 'Config':
        """Implement singleton pattern."""
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance._initialize()
        return cls._instance

    def _initialize(self) -> None:
        """Initialize configuration by loading environment variables."""
        load_dotenv()

        self.api_key = os.getenv('API_KEY')
        if not self.api_key:
            raise ValueError("API_KEY environment variable is required")

        self.log_level = os.getenv('LOG_LEVEL', 'DEBUG').upper()

        valid_levels = ['DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL']
        if self.log_level not in valid_levels:
            raise ValueError(f"Invalid LOG_LEVEL. Must be one of: {', '.join(valid_levels)}")

        logging.basicConfig(
            level=getattr(logging, self.log_level),
            format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
        )

    @classmethod
    def get_instance(cls) -> 'Config':
        return cls()

# Create a global instance
config = Config.get_instance()
