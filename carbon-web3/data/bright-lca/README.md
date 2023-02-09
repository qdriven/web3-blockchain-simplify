# Carbon-Footprint-intermodal

This is the code used in the analysis presented in:

_Pizzol M, Deterministic and stochastic carbon footprint of intermodal ferry andtruck freight transport across Scandinavian routes, Journal of Cleaner Production (2019),_ [doi.org/10.1016/j.jclepro.2019.03.270](https://doi.org/10.1016/j.jclepro.2019.03.270).

The supplementary materials to the article include the raw data: 

- inventory tables: `Ferries_ei4.csv` and `Routes_ei4.csv`
- LCIA result table `Static_contribution_analysis.csv`
- results of the Monte Carlo simulation `MC_simulation_1000_iter.csv`

These are needed to run the scripts contained in this repository, for example if you want to reproduce the results of the paper. If you don't have access to the article or the files, you can get them from me directly, please send a request to [massimo@plan.aau.dk](mailto:massimo@plan.aau.dk)


##The repository includes:

`HH_LCI.py` Python script to reproduce results of the LCA using the brightway2 LCA software. Imports the inventory in, performs LCA calculations and contribution analysis, then exports results.

`HH_MC.py` Python script to reproduce the Monte Carlo simulation using the brightway2 LCA software. Imports inventory in, runs comparative Monte Carlo, then export results.

`HH_Plots.R` R script used in generating the plots of the paper. 
`HH_Stat_analysis.R` R script with the statistical analysis of Monte Carlo results, mainly generating descriptive statistics and doing pairwise testing of different alternatives.`lci_to_bw2.py` A little python script to import inventory tables in .csv directly into brightway2.

**Please get in touch** if you find any mistake or problem in running these scripts.

##DOI and versions

Version identifier:
[![DOI](https://zenodo.org/badge/156857727.svg)](https://zenodo.org/badge/latestdoi/156857727)

You can cite all versions by using the DOI [10.5281/zenodo.1481760](https://doi.org/10.5281/zenodo.1481760). This DOI represents all versions, and will always resolve to the latest one. 

# Cite as

Pleas cite the specific version, for example for version 1.0:

_Massimo Pizzol. (2018, November 9). massimopizzol/Carbon-Footprint-intermodal: First release (Version 1.0). Zenodo. http://doi.org/10.5281/zenodo.1481761_


