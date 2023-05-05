# RAM dev instance process

1. Go to CFME
   https://calo-bgl-cfme.cisco.com/dashboard/show#/

2. Provision VM with the following details:

Service: "BGL VMware Provision from ISO"
Service Information: - Service Name: <NAME> - Service Description: <DESC> - Environment: Team (12 months)

VM information: - Provider: bgl-gen-vc.cisco.com - VM Name: <NAME> - Guest OS Type: Windows Server 2022 / CentOS7/8 (User perference) - ISO - Disk Size: 300 - Network: PS_BGL14-GEN-VLAN-2525 (Find an IP that is available in this subnet 10.127.251.X/255.255.255.224 Gateway: 10.127.251.1, if not use some other network in this bgl-gen-vc) - NIC: E1000 - CPUs - 4/8/12 - Memory - 16

3. Once VM is provisioned, configure the network settings:

   - IP: Free IP from VLAN selected
   - Gateway: Gateway of the VLAN selected
   - Primary DNS: 72.163.128.140
   - Secondary DNS: 64.104.128.236

4. Once network is setup and internet access is available, install the following:

   - NodeJS: 14.16.0
   - NPM: 8.19.1

5. Git pull the RAM repo and create a branch off of "STAGING":

   - https://wwwin-github.cisco.com/abhism/RAM.git

6. Install requirements:

   - cd to Client/
     - npm install
   - cd to Server/
     - pip install -r requirements

7. Run FE/BE:
   - cd Client/
     - npm start
   - cd Server/
     - uvicorn app.main:app --host 0.0.0.0 --port 8050 --reload
