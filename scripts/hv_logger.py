import logging                                                                                                                                  
import logging.config
 
LOG_FILE = './hv_auto_chart.log'
 
def setupLogger():
    logging_config = { 
        "version": 1,
        "formatters": {
            "standard": {
                "format": "%(asctime)s - %(levelname)s - %(name)s - %(message)s",
                "datefmt": "%Y-%m-%d %H:%M:%S",  # Custom date format
            },  
        },  
        "handlers": {
            "console": {
                "class": "logging.StreamHandler",
                "formatter": "standard",
                "level": "DEBUG",
            },  
            "file": {
                "class": "logging.FileHandler",
                "formatter": "standard",
                "level": "INFO",
                "filename": LOG_FILE,
            },  
        },  
        "root": {
            "handlers": ["console", "file"],
            "level": "DEBUG",
        },  
        "loggers": {
            "autochart": {
                "handlers": ["console", "file"],
                "level": "DEBUG",
                "propagate": False,
            },  
        },  
    }   
    logging.config.dictConfig(logging_config)
