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


projects
projects.set_current('HH2') # Create new project


# =============================================================================
# bw2setup() 
# 
# fpei34  = "/Users/massimo/Documents/AAU/Research/Databases/ecoinvent v3.4/datasets"
# 
# if 'ecoinvent 3.4 conseq' in databases:
#     print("Database has already been imported")
# else:
#     ei34 = SingleOutputEcospold2Importer(fpei34, 'ecoinvent 3.4 conseq')
#     ei34.apply_strategies()
#     ei34.statistics()
# 
# ei34.write_database()
# =============================================================================


databases #OK There is ecoinvent 3.4. in the project

if 'Routes' in databases: del databases['Routes']
if 'Ferries' in databases: del databases['Ferries']


# Import ferries
fer_data = pd.read_csv('Ferries_ei4.csv', header = 0, sep = ";", encoding = 'utf-8-sig') # important to specify encoding

# clean up 
fer_data = fer_data.drop(['Simapro names', 'BW2 names'], 1)  # remove the columns not needed
fer_data['Exchange uncertainty type'] = fer_data['Exchange uncertainty type'].fillna(0).astype(int) # uncertainty as integer 
                    ### Note: (can't have the full column if there are mixed nan and values, so use zero as default)
fer_data.head()    
fer_data.tail()   
fer_data.iloc[:,6] # no encoding problems

# Create a dict that can be written as database
fer_dict = lci_to_bw2(fer_data) # Perfect.
fer_dict

# Write a bw2 database
databases
if 'Ferries' in databases: del databases['Ferries']
ferries = Database("Ferries")
ferries.write(fer_dict)
[print(act) for act in ferries]


# Import routes
rou_data = pd.read_csv('Routes_ei4.csv', header = 0, sep = ";", encoding = 'utf-8-sig') # important to specify encoding
rou_data.head()
rou_data_routes = rou_data.loc[:,['Activity code','Country','From','To','Via','By', 'Ferry']].drop_duplicates()

# clean up
rou_data = rou_data.drop(['Simapro names', 'BW2 names', 'Country','From','To','Via','By','Ferry'], 1)  # remove the columns not needed
rou_data['Exchange uncertainty type'] = rou_data['Exchange uncertainty type'].fillna(0).astype(int) # uncertainty as integer 
                    ### Note: (can't have the full column if there are mixed nan and values, so use zero as default)
rou_data.head()    

rou_data_road = rou_data.loc[rou_data['Exchange database'] != 'Ferries']
rou_data_sea = rou_data.loc[rou_data['Exchange input'] != '92fc17f49413aa01c991319132f8788f']


mymethod = ('IPCC 2013', 'climate change', 'GWP 100a')


def dolcacalc(myact, demand):
    functional_unit = {myact: demand} 
    lca = LCA(functional_unit, mymethod)
    lca.lci()
    lca.lcia()
    return lca.score

def getLCAresults(mydb):
    bw2_db = lci_to_bw2(mydb) # Perfect.
    if 'Routes' in databases: del databases['Routes']
    t_db = Database("Routes")
    t_db.write(bw2_db)
    
    all_activities = []
    results = []
    for act in t_db:
        all_activities.append(act['name'])
        results.append(dolcacalc(act,1000))
        print(act['name'])
     
    results_dict = dict(zip(all_activities, results))
    #results_df = pd.DataFrame({'Route': all_activities, 'GWP': results})
    #results_df = results_df.sort_values(by =['Route'])

    return results_dict


results_road = getLCAresults(rou_data_road) # total impact per road
results_sea = getLCAresults(rou_data_sea) # total impact per sea
results_route = getLCAresults(rou_data) # total impact per route


# Some formatting

results_all = pd.DataFrame([results_road, results_sea]).T
results_all.columns = ['Road','Sea']
rou_data_routes = rou_data_routes.set_index('Activity code')
results_all = pd.merge(results_all, rou_data_routes, left_index=True, right_index=True)
results_all = results_all.sort_index(axis = 0, ascending=True)
results_all['FromToVia'] = results_all['From'].astype(str) + "-" + results_all['To'].astype(str) + " "+ results_all['Via'].astype(str)
results_all.to_csv('Static_contribution_analysis' + when + '.csv')


# check that the result is correct
results_rou = pd.DataFrame([results_route]).T
results_tot = results_all.Sea + results_all.Road
plt.plot(results_rou, results_tot) # alright


# Plotting

plt.figure(figsize=(40,40))
plt.style.use('ggplot')
results_all.plot(kind='barh', x = results_all.index, stacked=True)
plt.xlabel(methods[mymethod]["unit"]) # need to make this work soon or later
plt.ylabel("")
#plt.tight_layout()
plt.savefig('Contribution_analysis_detail_' + when + '.png')



# compare 1 tonkm lorry with 1 km ferry
demand = 1 # tonkm
mymethod = ('IPCC 2013', 'climate change', 'GWP 100a')

lorry = Database("ecoinvent 3.4 conseq").get('92fc17f49413aa01c991319132f8788f') # lorry

functional_unit = {lorry: demand} 
lca = LCA(functional_unit, mymethod)
lca.lci()
lca.lcia()
lorry_score = lca.score
lorry_res = pd.DataFrame([[lorry['name'], lorry['unit'], lorry_score],], columns = ['1 unit of','Unit','kg CO2-eq'])

print(lorry['name'], lorry['unit'])

fers = []

for fer in ferries:
    #if fer['name'] != ('LNG combustion') and fer['name'] != 'Diesel combustion':
        functional_unit = {fer: demand} 
        lca = LCA(functional_unit, mymethod)
        lca.lci()
        lca.lcia()
        print(fer)
        print(lca.score)  # OK
        fers.append({'1 unit of' : fer['name'], 'Unit' : fer['unit'], 'kg CO2-eq' : lca.score})

df = pd.DataFrame(fers)
df = df.append(lorry_res)
print(df)
df.to_csv('comparison_' + when + '.csv')

