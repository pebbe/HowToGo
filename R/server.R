# install.packages("corrplot")
# install.packages("Rserve")

library("corrplot")
library("Rserve")

generateCorrelationPlot <- function () {
  # create temporary png file
  filename <- tempfile("plot", fileext = ".png")
  png(filename)
  
  # draw graph to png image
  M <- cor(mtcars)
  corrplot(M)
  dev.off()

  # obtain image binary array
  image <- readBin(filename, "raw", .Machine$integer.max)
  unlink(filename)

  # return binary array to Go
  return (image)
}
run.Rserve()
