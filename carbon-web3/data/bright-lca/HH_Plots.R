library('ggplot2')
library('reshape2')

# Contribution analysis plot
getwd()
routes_wide = read.csv('Static_contribution_analysis.csv', sep = ',')
head(routes_wide, 30)
names(routes_wide)
summary(routes_wide)

routes_wide$Via <- gsub('via "storebælt" & "øresund" ', 'via Storebælt-Øresund', routes_wide$Via)
routes_wide$FromToVia <- gsub('via "storebælt" & "øresund "', 'via Storebælt-Øresund', routes_wide$FromToVia)

routes_wide['FromTo'] <- do.call(paste, c(routes_wide['From'],routes_wide['To'],sep = "-"))
routes_wide$tot = routes_wide$Road + routes_wide$Sea
routes <- melt(routes_wide, id.vars = c('X','Via', 'FromTo', 'tot'), measure.vars = c("Road", "Sea"))
head(routes, 30)

routes$FromTo <- factor(routes$FromTo, levels= c("Duisburg-Bergen", "Duisburg-Larvik", 
                                       "Duisburg-Oslo","Reims-Stavanger",
                                       "Reims-Kongsberg", "Reims-Oslo",
                                       "Hitra-Saint Laurent Blangy", "Hitra-Utska"))


routes$X <- gsub('Route', '', routes$X)
routes$Via <- gsub('via ', '', routes$Via)

g <- ggplot(routes, aes(x = X, y = value, fill = variable))+
  geom_bar(stat="identity")+
  #ggtitle("Figure 2. Global Warming Potential per route (kg CO2-eq)")+
  #coord_flip() +
  #geom_text(aes(label=round(value, digits = 0)), position = position_stack(vjust = 0.5), size=2.5)+
  theme_minimal()+
  theme(text = element_text(size = 10),
        #axis.title.x = element_text(size = 10),
        #axis.title.y = element_text(size = 10),
        #axis.title.y = element_blank(),
        #axis.title.x = element_blank(),
        axis.text.x=element_text(angle = 90, vjust = 0.5),
        #axis.text.y=element_blank(),
        #axis.ticks.y = element_line(),
        #axis.ticks.x = element_line(),
        #legend.position="top", 
        #legend.direction="horizontal", 
        #legend.title = element_blank(),
        legend.position = "none",
        panel.grid = element_blank(),
        panel.background = element_rect(fill = "white"),
        plot.margin = margin(0.5,0.5,0.5,0.5, "cm"))+
  stat_summary(fun.y = sum, aes(label = signif(..y.., digits = 5), group = X), geom = "text", angle = 90, hjust = 0.01, size=2.25)+
  facet_wrap(~FromTo, nrow = 2, scales = 'free_x')
g + xlab("Route number") + + ylab(expression("kg CO"[2]*"-eq")) + coord_cartesian(ylim = c(0, 275)) + scale_y_continuous(breaks = c(0,50,100,150,200,250))





# Monte Carlo plot
mc = read.csv('MC_simulation_1000_iter.csv', sep = ',')
#mc <- mc[,-1] # remove pandas index...
mc <- mc[,sort(names(mc))]
head(mc)

mc_long <- melt(mc, id = c('X'))
mc_long <- mc_long[,-1]
mc_long$variable <- as.character(mc_long$variable)
mc_long <- mc_long[order(mc_long$variable),]
head(mc_long)
tail(mc_long)


corridors <- c("Duisburg-Bergen", "Duisburg-Larvik","Duisburg-Oslo",
               "Reims-Stavanger", "Reims-Kongsberg", "Reims-Oslo",
               "Hitra-Saint Laurent Blangy", "Hitra-Utska")
routesnr <- c(8, 7, 8, 8, 7, 8, 10, 10)

cor_col <- c()

for (i in 1:length(corridors)){
  cor_i <- rep(corridors[i], 1000*routesnr[i])
  cor_col <- append(cor_col,cor_i)
  }

length(cor_col)

mc_long$FromTo <- cor_col
mc_long$variable <- factor(mc_long$variable)
mc_long$FromTo<- factor(mc_long$FromTo)
head(mc_long)
tail(mc_long)
levels(mc_long$FromTo)
levels(mc_long$variable)
mc_long$FromTo <- factor(mc_long$FromTo, levels= c("Duisburg-Bergen", "Duisburg-Larvik", 
                                                 "Duisburg-Oslo","Reims-Stavanger",
                                                 "Reims-Kongsberg", "Reims-Oslo",
                                                 "Hitra-Saint Laurent Blangy", "Hitra-Utska"))


mc_long$variable <- gsub('Route', '', mc_long$variable)

gmc <- ggplot(mc_long, aes(x = variable, y = value)) + 
  #geom_point()+
  geom_jitter(color = "red", width = 0.1, alpha = 0.15)+
  #geom_boxplot(alpha = 0.01, outlier.shape = NA)+
  stat_summary(fun.y=median, geom="point", size = 1)+
  #stat_summary(fun.y=mean, geom="point", shape = 6)+
  theme_minimal()+
  theme(text = element_text(size = 10),
        axis.text.x=element_text(angle = 90, vjust = 0.5),
        legend.position = "none",
        panel.grid = element_blank(),
        panel.background = element_rect(fill = "white"),
        plot.margin = margin(0.5,0.5,0.5,0.5, "cm"))+
  facet_wrap(~FromTo, nrow = 2, scales = 'free_x')
gmc + xlab("Route number") + ylab(expression("kg CO"[2]*"-eq")) 


