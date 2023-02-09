#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
@author: massimo
"""


import pandas as pd
import numpy as np
from lci_to_bw2 import *
from brightway2 import *
from matplotlib import pyplot as plt
import time
when = time.strftime("%Y%m%d")


projects.set_current('HH2')
databases

if 'Routes' in databases: del databases['Routes']
if 'Ferries' in databases: del databases['Ferries']
if 'HH' in databases: del databases['HH']


# Import ferries
fer_data = pd.read_csv('Ferries_ei4.csv', header = 0, sep = ";", encoding = 'utf-8-sig')
fer_data = fer_data.drop(['Simapro names', 'BW2 names'], 1)
fer_data['Exchange uncertainty type'] = fer_data['Exchange uncertainty type'].fillna(0).astype(int)  
fer_dict = lci_to_bw2(fer_data)
if 'Ferries' in databases: del databases['Ferries']
ferries = Database("Ferries")
ferries.write(fer_dict)
[print(act) for act in ferries]


# Import routes
rou_data = pd.read_csv('Routes_ei4.csv', header = 0, sep = ";", encoding = 'utf-8-sig')
rou_data_routes = rou_data.loc[:,['Activity code','Country','From','To','Via','By', 'Ferry']].drop_duplicates()
rou_data = rou_data.drop(['Simapro names', 'BW2 names', 'Country','From','To','Via','By','Ferry'], 1)  
rou_data['Exchange uncertainty type'] = rou_data['Exchange uncertainty type'].fillna(0).astype(int) 
rou_dict = lci_to_bw2(rou_data)  
routes = Database("Routes")
routes.write(rou_dict)
[print(act) for act in routes]

mymethod = ('IPCC 2013', 'climate change', 'GWP 100a')

for act in routes:
    functional_unit = {act: 100} 
    lca = LCA(functional_unit, mymethod)
    lca.lci()
    lca.lcia()


acts = []
demands = []

for act in routes:
    acts.append(act['name'])
    demands.append({act: 1000})


mc = MonteCarloLCA(demands[0], mymethod)
next(mc)

mc.redo_lcia(demands[0])
mc.score
mc.redo_lcia(demands[1])
mc.score

iterations = 1000
simulations = []

for _ in range(iterations):
    print(_)
    next(mc)
    mcresults = []    
    for i in demands:
        mc.redo_lcia(i)
        mcresults.append(mc.score)
    simulations.append(mcresults)
    

simulations
mcresults 
df = pd.DataFrame(simulations, columns = acts)
df.to_csv('MC_simulation_1000_iter'+when+'.csv')

df.plot(kind = 'box')
df.T.melt()
plt.hist(df.Route30.values)
df.Route20