import logging
import os
from datetime import datetime

def setupLogger():
    # Get the logger
    logger = logging.getLogger("autochart")
    logger.setLevel(logging.INFO)
    
    # Create a file handler for the logger
    log_filename = f"autochart_{datetime.now().strftime('%Y%m%d_%H%M%S')}.log"
    log_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'logs', log_filename)
    
    file_handler = logging.FileHandler(log_path)
    file_handler.setLevel(logging.INFO)
    
    # Create a console handler
    console_handler = logging.StreamHandler()
    console_handler.setLevel(logging.INFO)
    
    # Create a formatter and attach it to handlers
    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    file_handler.setFormatter(formatter)
    console_handler.setFormatter(formatter)
    
    # Add the handlers to the logger
    logger.addHandler(file_handler)
    logger.addHandler(console_handler)
    
    return logger
