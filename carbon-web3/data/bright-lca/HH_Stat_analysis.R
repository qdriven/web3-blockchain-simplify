# Statistical analysis of MC simulation results
# Massimo Pizzol, Feb 2018

library("reshape2")
library('broom')

getwd()

# Import & format data
mc = read.csv('MC_simulation_1000_iter.csv', sep = ',')
mc <- mc[,-1] # remove pandas index...
mc <- mc[,sort(names(mc))]
head(mc)
names(mc)
summary(mc)

# Export basic stats for reporting
stats_base <- t(sapply(mc,summary))
stats_sd <- sapply(mc,sd)
stats_nice <- cbind(stats_base, stats_sd)
write.table(stats_nice, "MC_stats.txt")


# Testing for normality.
for (i in 1:dim(mc)[2]){
  #hist(log(mcs[,i]), breaks = 100)
  print(names(mc)[i])
  print(shapiro.test(log(mc[,i])))
} # (almost) nothing normal here


# Some tests and plots

#stripcharts
par(mfrow=c(2,4), oma=c(0,0,2,0)) # nice and easy, found here: https://stat.ethz.ch/pipermail/r-help/2008-August/170110.html
stripchart(mc[,1:8], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Duisburg-Bergen', las = 2, vertical = TRUE)
stripchart(mc[,9:15], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Duisburg-Larvik', las = 2, vertical = TRUE)
stripchart(mc[,16:23], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Duisburg-Oslo', las = 2, vertical = TRUE)
stripchart(mc[,24:31], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Reims-Stavanger', las = 2, vertical = TRUE)
stripchart(mc[,32:38], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Reims-Kongsberg', las = 2, vertical = TRUE)
stripchart(mc[,39:46], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Reims-Oslo', las = 2, vertical = TRUE)
stripchart(mc[,47:56], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Hitra-Saint Laurent Blangy', las = 2, vertical = TRUE)
stripchart(mc[,57:66], pch = 16, method = 'jitter', jitter = 0.2, col = 'red', main = 'Hitra-Utska', las = 2, vertical = TRUE)
title(main = "Figure 3. Carbon Footprint: results of Monte Carlo simulation (kg CO2-eq)", outer = TRUE)


boxplot(log(mc[,1:8]), col = 'lightblue', las = 2, main = 'Duisburg-Bergen')
boxplot(log(mc[,9:15]), col = 'lightblue', las = 2, main = 'Duisburg-Larvik')
boxplot(log(mc[,16:23]), col = 'lightblue', las = 2, main = 'Duisburg-Oslo')
boxplot(log(mc[,24:31]), col = 'lightblue', las = 2, main = 'Reims-Stavanger')
boxplot(log(mc[,32:38]), col = 'lightblue', las = 2, main = 'Reims-Kongsberg')
boxplot(log(mc[,39:46]), col = 'lightblue', las = 2, main = 'Reims-Oslo')
boxplot(log(mc[,47:56]), col = 'lightblue', las = 2, main = 'Hitra-Saint Laurent Blangy')
boxplot(log(mc[,57:66]), col = 'lightblue', las = 2, main = 'Hitra-Utska')

?stripchart
par(mfrow=c(1,1))

wilcox.test(mc$Route01,mc$Route04, paired = FALSE) # it's not paired (paired is before/after)
wilcox.test(log(mc$Route01),log(mc$Route04), paired = FALSE)

t.test(mc$Route01,mc$Route04, paired = FALSE) # it's not paired (paired is before/after)
t.test(log(mc$Route01),log(mc$Route04), paired = FALSE)

d <- mc[,1:8]  # reshape data in long format to perform pairwise tests
d <- t(d)
d <- cbind(rownames(d), data.frame(d, row.names=NULL))
colnames(d)[1] <- 'route'
d <- melt(d, id = 'route')

pairwise.wilcox.test(d$value, d$route, p.adj = "bonf")  # pairwise test with bonf correction
pairwise.t.test(d$value, d$route, p.adj = "bonf")  # not sure this is correct.


# Test all corridor routes pairwise for all corridors
corrid <- list(mc[,1:8],  # Duisburg	Bergen
               mc[,9:15],  # Duisburg	Larvik
               mc[,16:23],  # Duisburg Oslo
               mc[,24:31],  # Reims	Stavanger 
               mc[,32:38],  # Reims	Kongsberg 
               mc[,39:46],  # Reims	Oslo
               mc[,47:56],  # Hitra	Saint Laurent Blangy
               mc[,57:66])  # Hitra	Utska


wil_res <- list()


for (i in 1:length(corrid)){
  d <- t(corrid[[i]])
  d <- cbind(rownames(d), data.frame(d, row.names=NULL))
  colnames(d)[1] <- 'route'
  d <- melt(d, id = 'route')
  
  # wilcox test with bonf correction
  res <- pairwise.wilcox.test(d$value, d$route, p.adj = "bonf")  
  # Note, this is exactly the same as using wilcox.test(paired = TRUE) but allows using the correction.
  #clean = data.frame(apply(res$p.value, 2, format.pval))
  wil_res[[i]] <- res
}

?pairwise.wilcox.test

sink('MC_analysis_pairwise_wilcoxon.txt')  # Export results in file
for (i in wil_res){
  print(tidy(i))
}
sink()


# Calculate percentages of differneces between routes, examples
A <- mc$Route01
B <- mc$Route04
par(mfrow=c(1,1))
hist(A-B, breaks = 100) # as Simapro does
sum((A-B) > 0) / 1000 *100 # perc positive difference (B better than A)
sum((A-B) < 0) / 1000 *100 # perc negative difference (A better than B)
sum((A-B) == 0) / 1000 *100 # perc positive difference (A equal to B)


# Do this for all route pairs in all corridors
diff_res_all <- list()

for (n in 1:length(corrid)){
  
  cr <- corrid[[n]]  # cr = 'corridor route'
  diff_names <- list() 
  diff_values <- list()
  diff_perc <- list()
  
  for (i in 1:dim(cr)[2]){
    for (j in 1:dim(cr)[2]){
    
      names <- c(names(cr)[i], names(cr)[j])
      diff = cr[,i] - cr[,j]
      values <- tidy(summary(diff))
      perpos <- sum(diff > 0) / 1000 * 100 
      perneg <- sum(diff < 0) / 1000 * 100 
      perequ <- sum(diff == 0) / 1000 * 100
      perscal <- round(median(cr[,i]) / median(cr[,j]) * 100, 2) # impact of i compared to j
      percs <- c(perpos, perneg, perequ, perscal)
      
      index <- (i-1)*dim(cr)[2]+j
      diff_names[[index]] <- names
      diff_values[[index]] <- values
      diff_perc[[index]] <- percs
      
      }
    }
  
  doubles <- c() # this is an index to locate duplicate couples (e.g.Route01-Route02, Route02-Route01, I keep only the first)
  for (i in 1:dim(cr)[2]){  
    for (j in 1:i){
      doubles <- append(doubles,dim(cr)[2]*i-dim(cr)[2]+j)  # this indexing was clever :)
      }
    }
  
  diff_res <- cbind(do.call(rbind, diff_names), do.call(rbind, diff_values), do.call(rbind, diff_perc))
  names(diff_res)[1:2] <- c('BaseRoute','AltRoute')
  names(diff_res)[9:12] <- c('Percpos', 'Percneg', 'Percequ', 'Percscal')
  diff_res <- diff_res[-doubles,]  # removing duplicates
  
  diff_res_all[[n]] <- diff_res
  
}

diff_res_all_df <- do.call(rbind, diff_res_all)  
tail(diff_res_all_df, 10)  #just a check
write.table(diff_res_all_df, "MC_analysis_perc_diff.txt")
