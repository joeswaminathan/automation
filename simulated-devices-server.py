from aiohttp import web
import logging
import json
import traceback
import uuid
import os
from datetime import datetime

# Create logs directory before initializing logger
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
LOG_DIR = os.path.join(SCRIPT_DIR, 'logs')
# Create the logs directory if it doesn't exist
os.makedirs(LOG_DIR, exist_ok=True)

# Check if hv_logger.py exists, create it if not
HV_LOGGER_PATH = os.path.join(SCRIPT_DIR, 'hv_logger.py')
if not os.path.exists(HV_LOGGER_PATH):
    with open(HV_LOGGER_PATH, 'w') as f:
        f.write("""import logging
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
""")
    print(f"Created hv_logger.py at {HV_LOGGER_PATH}")

# Import the logger after ensuring it exists
from scripts.hv_logger import setupLogger

# Initialize logger
setupLogger()
logger = logging.getLogger("autochart")

# In-memory storage for our simulated devices
simulated_cameras = {}
simulated_monitors = {}

# Request logging middleware
@web.middleware
async def request_logger_middleware(request, handler):
    request_id = str(uuid.uuid4())[:8]
    start_time = datetime.now()
    
    # Log the incoming request details
    logger.info(f"[{request_id}] Request received: {request.method} {request.path} - Client: {request.remote}")
    
    if request.content_type == 'application/json':
        try:
            body = await request.json()
            logger.info(f"[{request_id}] Request body: {json.dumps(body)}")
        except:
            logger.info(f"[{request_id}] Could not parse JSON body")
    
    # Process the request
    try:
        response = await handler(request)
        
        # Log the response
        status_code = response.status
        duration = (datetime.now() - start_time).total_seconds()
        logger.info(f"[{request_id}] Response sent: {status_code} - Took {duration:.4f}s")
        
        return response
    except Exception as e:
        logger.error(f"[{request_id}] Error processing request: {str(e)}")
        logger.error(traceback.format_exc())
        duration = (datetime.now() - start_time).total_seconds()
        logger.info(f"[{request_id}] Response failed - Took {duration:.4f}s")
        raise

# --- Camera Endpoints ---

async def get_all_cameras(request):
    """Get a list of all simulated cameras using POST"""
    try:
        # You can add filter parameters in the request body
        filter_params = await request.json() if request.body_exists else {}
        logger.info(f"Returning all simulated cameras with filters: {filter_params}")
        
        # Apply any filters here if needed
        cameras = list(simulated_cameras.values())
        
        return web.json_response({"cameras": cameras})
    except Exception as e:
        logger.error(f"Error retrieving cameras: {str(e)}")
        return web.json_response({"error": "Failed to retrieve cameras"}, status=500)

async def get_camera(request):
    """Get a specific camera by ID using GET"""
    camera_id = request.match_info.get('camera_id')
    logger.info(f"Request for camera with ID: {camera_id}")
    
    if camera_id in simulated_cameras:
        return web.json_response(simulated_cameras[camera_id])
    else:
        return web.json_response({"error": "Camera not found"}, status=404)

async def create_camera(request):
    """Create a new simulated camera"""
    try:
        camera_data = await request.json()
        
        # Generate a new ID if not provided
        if 'id' not in camera_data:
            camera_data['id'] = str(uuid.uuid4())
            
        camera_id = camera_data['id']
        
        # Add creation timestamp
        camera_data['created_at'] = datetime.now().isoformat()
        camera_data['status'] = 'online'  # Default status
        
        # Store the camera
        simulated_cameras[camera_id] = camera_data
        
        logger.info(f"Created new camera with ID: {camera_id}")
        return web.json_response(camera_data, status=201)
    except Exception as e:
        logger.error(f"Error creating camera: {str(e)}")
        return web.json_response({"error": "Invalid camera data"}, status=400)

async def get_all_cameras_status(request):
    """Get status of all cameras using POST"""
    try:
        # You can add filter parameters in the request body
        filter_params = await request.json() if request.body_exists else {}
        logger.info(f"Returning status for all cameras with filters: {filter_params}")
        
        statuses = {}
        for camera_id, camera in simulated_cameras.items():
            statuses[camera_id] = {
                "id": camera_id,
                "status": camera.get('status', 'unknown'),
                "last_updated": datetime.now().isoformat()
            }
        
        logger.info(f"Returning status for {len(statuses)} cameras")
        return web.json_response({"camera_statuses": statuses})
    except Exception as e:
        logger.error(f"Error retrieving camera statuses: {str(e)}")
        return web.json_response({"error": "Failed to retrieve camera statuses"}, status=500)

async def get_camera_status(request):
    """Get status of a specific camera using GET"""
    camera_id = request.match_info.get('camera_id')
    logger.info(f"Request for status of camera with ID: {camera_id}")
    
    if camera_id in simulated_cameras:
        status = {
            "id": camera_id,
            "status": simulated_cameras[camera_id].get('status', 'unknown'),
            "last_updated": datetime.now().isoformat()
        }
        return web.json_response(status)
    else:
        return web.json_response({"error": "Camera not found"}, status=404)

# --- Multi-Vital Monitor Endpoints ---

async def get_all_monitors(request):
    """Get a list of all simulated vital monitors using POST"""
    try:
        # You can add filter parameters in the request body
        filter_params = await request.json() if request.body_exists else {}
        logger.info(f"Returning all simulated multi-vital monitors with filters: {filter_params}")
        
        # Apply any filters here if needed
        monitors = list(simulated_monitors.values())
        
        return web.json_response({"monitors": monitors})
    except Exception as e:
        logger.error(f"Error retrieving monitors: {str(e)}")
        return web.json_response({"error": "Failed to retrieve monitors"}, status=500)

async def get_monitor(request):
    """Get a specific vital monitor by ID using GET"""
    monitor_id = request.match_info.get('monitor_id')
    logger.info(f"Request for monitor with ID: {monitor_id}")
    
    if monitor_id in simulated_monitors:
        return web.json_response(simulated_monitors[monitor_id])
    else:
        return web.json_response({"error": "Monitor not found"}, status=404)

async def create_monitor(request):
    """Create a new simulated vital monitor"""
    try:
        monitor_data = await request.json()
        
        # Generate a new ID if not provided
        if 'id' not in monitor_data:
            monitor_data['id'] = str(uuid.uuid4())
            
        monitor_id = monitor_data['id']
        
        # Add creation timestamp and default status
        monitor_data['created_at'] = datetime.now().isoformat()
        monitor_data['status'] = 'online'  # Default status
        
        # Store the monitor
        simulated_monitors[monitor_id] = monitor_data
        
        logger.info(f"Created new multi-vital monitor with ID: {monitor_id}")
        return web.json_response(monitor_data, status=201)
    except Exception as e:
        logger.error(f"Error creating monitor: {str(e)}")
        return web.json_response({"error": "Invalid monitor data"}, status=400)

async def get_all_monitors_status(request):
    """Get status of all vital monitors using POST"""
    try:
        # You can add filter parameters in the request body
        filter_params = await request.json() if request.body_exists else {}
        logger.info(f"Returning status for all monitors with filters: {filter_params}")
        
        statuses = {}
        for monitor_id, monitor in simulated_monitors.items():
            statuses[monitor_id] = {
                "id": monitor_id,
                "status": monitor.get('status', 'unknown'),
                "last_updated": datetime.now().isoformat(),
                "vital_signs": monitor.get('vital_signs', {})
            }
        
        logger.info(f"Returning status for {len(statuses)} monitors")
        return web.json_response({"monitor_statuses": statuses})
    except Exception as e:
        logger.error(f"Error retrieving monitor statuses: {str(e)}")
        return web.json_response({"error": "Failed to retrieve monitor statuses"}, status=500)

async def get_monitor_status(request):
    """Get status of a specific vital monitor using GET"""
    monitor_id = request.match_info.get('monitor_id')
    logger.info(f"Request for status of monitor with ID: {monitor_id}")
    
    if monitor_id in simulated_monitors:
        status = {
            "id": monitor_id,
            "status": simulated_monitors[monitor_id].get('status', 'unknown'),
            "last_updated": datetime.now().isoformat(),
            "vital_signs": simulated_monitors[monitor_id].get('vital_signs', {})
        }
        return web.json_response(status)
    else:
        return web.json_response({"error": "Monitor not found"}, status=404)

# --- Device Management Endpoints ---

async def handle_device_command(request):
    """Handle device management commands"""
    try:
        command_data = await request.json()
        command = command_data.get('command', '')
        logger.info(f"Received device command: {command}")
        
        if command == 'add':
            device_type = command_data.get('type', '')
            device_id = command_data.get('id', str(uuid.uuid4()))
            
            if device_type == 'camera':
                simulated_cameras[device_id] = {
                    'id': device_id,
                    'type': 'camera',
                    'status': 'online',
                    'created_at': datetime.now().isoformat(),
                    **command_data
                }
                logger.info(f"Added simulated camera with ID: {device_id}")
                return web.json_response({"status": "success", "id": device_id})
                
            elif device_type == 'monitor':
                simulated_monitors[device_id] = {
                    'id': device_id,
                    'type': 'monitor',
                    'status': 'online',
                    'created_at': datetime.now().isoformat(),
                    **command_data
                }
                logger.info(f"Added simulated monitor with ID: {device_id}")
                return web.json_response({"status": "success", "id": device_id})
            
            else:
                return web.json_response({"error": f"Unknown device type: {device_type}"}, status=400)
                
        elif command == 'delete':
            device_id = command_data.get('id', '')
            
            if device_id in simulated_cameras:
                del simulated_cameras[device_id]
                logger.info(f"Deleted simulated camera with ID: {device_id}")
                return web.json_response({"status": "success"})
                
            elif device_id in simulated_monitors:
                del simulated_monitors[device_id]
                logger.info(f"Deleted simulated monitor with ID: {device_id}")
                return web.json_response({"status": "success"})
                
            else:
                return web.json_response({"error": f"Device not found: {device_id}"}, status=404)
                
        else:
            return web.json_response({"error": f"Unknown command: {command}"}, status=400)
            
    except Exception as e:
        logger.error(f"Error processing device command: {str(e)}")
        return web.json_response({"error": "Invalid command data"}, status=400)

# --- SaaS Message Handler (from your original code) ---

async def handle_saas_message(message):
    # This is a stub for your existing functionality
    logger.info(f"Handling SaaS message: {message}")
    return {"status": "processed", "message_id": message.get('id', 'unknown')}

async def receive_saas_message(request):
    try:
        message = await request.json()
        batch = message.get('batch', [message])  # Fallback to single message
        if not batch:
            raise Exception("Invalid batch message")
        
        responses = [await handle_saas_message(msg) for msg in batch]
        response_msg = {"batch": responses}
        
        logger.info(f"Responding: {response_msg}")
        return web.json_response(response_msg)
    except Exception as e:
        logger.exception(f"Exception = {e}")
        logger.error(f"receive_saas_message: Exception = {e}")
        response = {"error": "Invalid message"}
        return web.json_response(response, status=400)

async def receive_response_message(request):
    # This is a stub for your existing functionality
    logger.info("Received response message")
    try:
        message = await request.json()
        logger.info(f"Response message content: {message}")
        return web.json_response({"status": "received"})
    except Exception as e:
        logger.error(f"Error processing response message: {str(e)}")
        return web.json_response({"error": "Invalid message"}, status=400)

# --- Main Application Setup ---

def create_app():
    # Create application with middleware for logging
    app = web.Application(middlewares=[request_logger_middleware])
    
    # Set up camera routes - using POST for multiple devices, GET for single device
    app.router.add_post('/devices/simulated/cameras', get_all_cameras)  # Changed from GET to POST
    app.router.add_post('/devices/simulated/cameras/new', create_camera)
    app.router.add_get('/devices/simulated/cameras/{camera_id}', get_camera)
    app.router.add_post('/devices/simulated/cameras/status', get_all_cameras_status)  # Changed from GET to POST
    app.router.add_get('/devices/simulated/cameras/status/{camera_id}', get_camera_status)
    
    # Set up multi-vital monitor routes - using POST for multiple devices, GET for single device
    app.router.add_post('/devices/simulated/multi-vital-monitors', get_all_monitors)  # Changed from GET to POST
    app.router.add_post('/devices/simulated/multi-vital-monitors/new', create_monitor)
    app.router.add_get('/devices/simulated/multi-vital-monitors/{monitor_id}', get_monitor)
    app.router.add_post('/devices/simulated/multi-vital-monitors/status', get_all_monitors_status)  # Changed from GET to POST
    app.router.add_get('/devices/simulated/multi-vital-monitors/status/{monitor_id}', get_monitor_status)
    
    # Add device management endpoint for simple integration with original code
    app.router.add_post('/device', handle_device_command)
    
    # Add the existing routes from your original code
    app.router.add_post('/command', receive_saas_message)
    app.router.add_post('/status', receive_saas_message)
    app.router.add_post('/response', receive_response_message)
    
    return app

if __name__ == "__main__":
    app = create_app()
    logger.info("Starting simulated devices REST server on http://0.0.0.0:8081")
    web.run_app(app, host='0.0.0.0', port=8081)






    