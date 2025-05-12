1. Simulate video for all 11 patients
    a) Start
        sim_dev_test.py demo-sim-camera-add.json
    b) Stop
        sim_dev_test.py demo-sim-camera-delete.json
2. Simulate Vitals for 7 patients (except the four in the top row)
    a) Start
        sim_dev_test.py demo-sim-vital-7-patient-add.json
    b) Stop
        sim_dev_test.py demo-sim-vital-7-patient-delete.json
3. Simulate Vitals for Maverick and Clara
    a) Start
        sim_dev_test.py demo-sim-vital-2-patient-add.json
    b) Stop
        sim_dev_test.py demo-sim-vital-2-patient-delete.json
4. Simulate Vitals for Jameswell and Michael
    a) Start
        sim_dev_test.py demo-sim-vital-2a-patient-add.json
    b) Stop
        sim_dev_test.py demo-sim-vital-2a-patient-delete.json

