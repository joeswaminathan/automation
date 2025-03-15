#!/usr/bin/python3

import aiohttp
import asyncio
import logging
from hv_logger import setupLogger
import json
import os
import sys
from pprint import pprint
import threading
from concurrent.futures import ThreadPoolExecutor

setupLogger()
logger = logging.getLogger("autochart")

patientDB = {}
deviceDB = {}

class HttpOps(object):
    def __init__(self, url_base, headers=None):
        self.url = url_base
        self.headers = headers

    def makeUrl(self, endpoint):
        return self.url + endpoint

    async def get(self, url):
        async with aiohttp.ClientSession() as session:
            async with session.get(url, headers=self.headers) as response:
        # Check if the request was successful
                if response.status < 300:
                    data = await response.json()
                    #logger.info(f"Response: {data}")
                    return data
                else:
                    logger.debug(f"Get Request failed with status: {response}")
                    logger.info(f"Get Request failed with status: {response.status}")
        return None

    async def post(self, url, payload):
        async with aiohttp.ClientSession() as session:
            async with session.post(url, json=payload, headers=self.headers) as response:
                # Check if the request was successful
                if response.status < 300:
                    data = await response.json()
                    #logger.info(f"Response: {data}")
                    return data
                else:
                    logger.info(f"Post Request failed with status: {response.status}")
        return None

    async def delete(self, url):
        async with aiohttp.ClientSession() as session:
            async with session.delete(url, headers=self.headers) as response:
                # Check if the request was successful
                if response.status < 300:
                    data = await response.json()
                    #logger.info(f"Response: {data}")
                    return data
                else:
                    logger.info(f"Delete Request failed with status: {response.status}")
        return None


class SaaS(HttpOps):
    def __init__(self, token, url_base):
        super().__init__( url_base,
                          { "Authorization" : "Bearer " + token,
                            "Content-Type" : "application/json",
                            "Accept" : "application/json" } )   

    async def read_patient(self, payload):
        global patientDB, deviceDB
        response = await self.get_patient_by_id(payload['id'], { "id" : payload['uuid'] })
 
    async def read_device(self, payload):
        global patientDB, deviceDB
        response = await self.get_device_by_id(payload['id'], { "id" : payload['uuid'] })

    async def read_DB(self, payload):
        global patientDB, deviceDB
        with open(payload['JsonDB'], "r") as file:
            data = json.load(file)
        commands = data['commands']
        for c in commands:
                if c["command"] == "ReadPatient":
                    await self.read_patient(c["payload"])
                if c["command"] == "ReadDevice":
                    await self.read_device(c["payload"])

    async def add_patient(self, pId, payload):
        global patientDB, deviceDB
        url = self.makeUrl("patients/new")
        response = await self.post(url, payload)
        if response is not None:
            patientDB[pId] = response['data']
        return response['data']

    async def add_device(self, dId, payload):
        global patientDB, deviceDB
        url = self.makeUrl("devices/new")
        response = await self.post(url, payload)
        if response is not None:
            deviceDB[dId] = response['data']
        return response['data']

    async def delete_patient(self, pId):
        global patientDB, deviceDB
        patient = patientDB.get(pId, None)
        if patient is None:
            print(f"Patient ({pId}) doesn't exist")
            return
        url = self.makeUrl(f"patients/{patient['uuid']}")
        response = await self.delete(url)
        del patientDB[pId]
        return response

    async def delete_device(self, dId):
        global patientDB, deviceDB
        device = deviceDB.get(dId, None)
        if device is None:
            print(f"Patient ({dId}) doesn't exist")
            return
        url = self.makeUrl(f"devices/{device['uuid']}")
        response = await self.delete(url)
        del deviceDB[dId]
        return response

    async def get_device_by_id(self, dId, payload):
        global patientDB, deviceDB
        url = self.makeUrl(f"devices/{payload['id']}")
        filters = {"page": 1, "limit": 125, "sortBy": "type", "sortOrder": "DESC", "q": "", "type": "DEVICE"}
        response = await self.post(url, filters)
        if response is not None:
            deviceDB[dId] = response['data']
        return response['data']

    async def get_patient_by_id(self, pId, payload):
        global patientDB, deviceDB
        url = self.makeUrl(f"patients/{payload['id']}")
        response = await self.get(url)
        if response is not None:
            patientDB[pId] = response['data']
        return response['data']

    async def get_device(self, dId):
        global patientDB, deviceDB
        device = deviceDB.get(dId, None)
        if device is None:
            print(f"Device ({dId}) doesn't exist")
            return None
        device = await self.get_device_by_id(dId, { "id" : device['uuid'] })
        return device

    async def get_patient(self, pId):
        global patientDB, deviceDB
        patient = patientDB.get(pId, None)
        if patient is None:
            print(f"Patient ({pId}) doesn't exist")
            return None
        patient = await self.get_patient_by_id(pId, { "id" : patient['uuid'] })
        return patient

    async def get_patients(self):
        global patientDB, deviceDB
        url = self.makeUrl("patients")
        filters = {"page": 1, "limit": 125, "sortBy": "type", "sortOrder": "DESC", "q": "", "type": "DEVICE"}
        response = await self.post(url, filters)
        if response is not None:
            for patient in response['data']['items']:
                print(f"patient={patient['uuid']}: {patient['fname']} {patient['lname']}")
            return response['data']['items']
        else:
            return None

    async def get_devices(self):
        global patientDB, deviceDB
        url = self.makeUrl("devices")
        filters = {"page": 1, "limit": 125, "sortBy": "type", "sortOrder": "DESC", "q": "", "type": "DEVICE"}
        response = await self.post(url, filters)
        if response is not None:
            for device in response['data']['items']:
                print(f"device={device['uuid']}: {device['deviceSerial']}")
            return response['data']['items']
        else:
            return None

    async def get_patient_devices(self, pId):
        global patientDB, deviceDB
        patient = patientDB.get(pId, None)
        if patient is None:
            print(f"Patient ({pId}) doesn't exist")
            return None
        url = self.makeUrl(f"patients/{patient['uuid']}/devices")
        response = await self.get(url)
        devices = {}
        if response is not None:
            for device in response['data']['items']:
                devices[device['deviceSerial']] = device
                print(f"device={device['uuid']}: {device['deviceSerial']}")
            return devices
        return devices

    async def attach_device(self, pId, dId):
        global patientDB, deviceDB
        patient = patientDB.get(pId, None)
        # Refresh the stored data
        device = await self.get_device(dId)
        devices = await self.get_patient_devices(pId)
        attachedDevice = devices.get(device['deviceSerial'], None)
        if attachedDevice is not None:
            print(f"Error device ({dId}) is alread attached to patient ({pId})")
            return
        if not (device['patient'] is None or device['patient']['uuid'] is None):
            print(f"Error device is already attached to some patient ({device['patient']})")
            return
        url = self.makeUrl(f"patients/{patient['uuid']}/devices/new")
        payload = {
            "deviceUUID": device['uuid'],
            "mac": device["mac"],
            "deviceSerial": device["deviceSerial"],
            "configuration": 0.15,
            "type": device["type"],
            "branch": device["branch"]
        }
        response = await self.post(url, payload)
        if response is None:
            return None
        return response['data']

    async def detach_device(self, pId, dId):
        global patientDB, deviceDB
        patient = patientDB.get(pId, None)
        device = await self.get_device(dId)
        if device['patient'] is None:
            print(f"Error device ({dId}) is not attached to patient ({pId})")
            return None
        devices = await self.get_patient_devices(pId)
        attachedDevice = devices.get(device['deviceSerial'], None)
        if attachedDevice is None:
            print(f"Error device ({dId}) is not attached to patient ({pId})")
            return
        url = self.makeUrl(f"patients/{patient['uuid']}/devices/{attachedDevice['uuid']}")
        await self.delete(url)

    async def detach_all_devices_of_patient(self, pId):
        global patientDB, deviceDB
        patient = patientDB[pId]
        if patient is None:
            print(f"Error: patient ({pId}) doesn't exist")
            return
        devices = await self.get_patient_devices(pId)
        for device in devices.values():
            url = self.makeUrl(f"patients/{patient['uuid']}/devices/{device['uuid']}")
            await self.delete(url)
            print(f"Device ({device['deviceUUID']}) is detached from "
                  f"patient ({patient['uuid']})")

    async def detach_all_patients(self):
        global patientDB, deviceDB
        patients = await SD.get_patients()
        for idx, patient in enumerate(patients):
            pId = f"-detach-patient-{idx}"
            patientDB[pId] = patient
            await SD.detach_all_devices_of_patient(pId)
            del patientDB[pId]

    async def delete_all_patients(self):
        global patientDB, deviceDB
        await self.detach_all_patients()
        patients = await SD.get_patients()
        for idx, patient in enumerate(patients):
            pId = f"-delete-patient-{idx}"
            patientDB[pId] = patient
            await SD.delete_patient(pId)
        patientDB = {}

    async def delete_all_devices(self):
        global patientDB, deviceDB
        await self.detach_all_patients()
        devices = await SD.get_devices()
        for idx, device in enumerate(devices):
            dId = f"-delete-device-{idx}"
            deviceDB[dId] = device
            await SD.delete_device(dId)
        deviceDB = {}

    async def write_DB(self, payload):
        global patientDB, deviceDB
        outDB = {}
        commands = []
        for id in patientDB.keys():
            commands.append({"command" : "ReadPatient", "payload" : {"id" : id, "uuid" : patientDB[id]['uuid']}})
        for id in deviceDB.keys():
            commands.append({"command" : "ReadDevice", "payload" : {"id" : id, "uuid" : deviceDB[id]['uuid']}})
        outDB['commands'] = commands
        loop = asyncio.get_running_loop()
        with ThreadPoolExecutor() as pool:
            await loop.run_in_executor(pool, lambda: json.dump(outDB, open(payload['JsonDB'], "w"), indent=4))


simulatedDB ={}

class Edge(HttpOps):
    def __init__(self, url_base):
        super().__init__( url_base,
                          { "Content-Type" : "application/json",
                            "Accept" : "application/json" } )

    async def add_simulated_device(self, sId, payload):
        global simulatedDB
        url = self.makeUrl("device")
        response = await self.post(url, payload)
        if response is not None:
            simulatedDB[sId] = payload
            return True
        return False

    async def delete_simulated_device(self, sId, payload):
        global simulatedDB
        url = self.makeUrl("device")
        response = await self.post(url, payload)
        if response is not None:
            del simulatedDB[payload['id']]
            return True
        return False
        """
        payload = simulatedDB.get(sId, None)
        if payload is None:
            print(f"Error: Simulated device ({sId}) not found")
        payload['command'] = "delete"
        response = await self.post(url, payload)
        if response is None:
            print(f"Error: In deleting simulated device ({sId})")
            return False
        del simulatedDB[sId]
        return True
        """

    async def get_media_streams(self):
        url = self.makeUrl("v3/paths/list")
        response = await self.get(url)
        if response is None:
            return []
        streams = {}
        for stream in response['items']:
            streams[stream['name']] = stream
        return streams

ED = Edge("http://13.234.227.253:8118/")
SD = SaaS(os.getenv('TOKEN'),
          "https://in.livehealthyvibes.com/searchmanager/api/v1/")

async def execute_command(command):
    global ED, SD
    match command['command']:
        case "Pause":
            logger.info(f"Pausing for {command['payload']['duration']} seconds")
            await asyncio.sleep(command['payload']['duration'])
        case "ReadDB":
            await SD.read_DB(command['payload'])
        case "WriteDB":
            await SD.write_DB(command['payload'])
        case "GetPatients":
            await SD.get_patients()
        case "GetDevices":
            await SD.get_devices()
        case "ReadPatient":
            await SD.read_patient(command["payload"])
        case "ReadDevice":
            await SD.read_device(command["payload"])
        case "AddPatient":
            await SD.add_patient(command['id'],
                                 command['payload'])
        case "AddDevice":
            await SD.add_device(command['id'],
                                command['payload'])
        case "DeletePatient":
            await SD.delete_patient(command['id'])
        case "DeleteDevice":
            await SD.delete_device(command['id'])
        case "GetPatientById":
            await SD.get_patient_by_id(command['id'],
                                       command['payload'])
        case "GetDeviceById":
            await SD.get_device_by_id(command['id'],
                                      command['payload'])
        case "GetPatients":
            await SD.get_patients()
        case "GetDevices":
            await SD.get_devices()
        case "GetPatientDevices":
            await SD.get_patient_devices(command['id'])
        case "AttachDevice":
            await SD.attach_device(command['payload']['patient_id'],
                                   command['payload']['device_id'])
        case "DetachDevice":
            await SD.detach_device(command['payload']['patient_id'],
                                   command['payload']['device_id'])
        case "DeleteAllPatients":
            await SD.delete_all_patients()
        case "DeleteAllDevices":
            await SD.delete_all_devices()
        case "AddSimulatedDevice":
            await ED.add_simulated_device(command['id'],
                                          command['payload'])
        case "DeleteSimulatedDevice":
            await ED.delete_simulated_device(command['id'],
                                             command['payload'])
        case "GetMediaStreams":
            await ED.get_media_streams()
        case _ :
            print(f"Error: command ({command}) not recognized")

async def main():
    if len(sys.argv) < 2:
        print(f"Syntax : {sys.argv[0]} <test data file>")
    name = sys.argv[1]

    with open(name, 'r') as file:
        data = json.load(file)

    for command in data['commands']:
        await execute_command(command)

if __name__ == "__main__":
    asyncio.run(main())
