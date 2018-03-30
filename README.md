# jp

Dead simple terminal plots from JSON (or CSV) data. Bar charts, line charts, and scatter plots are supported.

[![Build Status](https://travis-ci.org/sgreben/jp.svg?branch=master)](https://travis-ci.org/sgreben/jp)

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
- [Examples](#examples)
    - [Bar chart](#bar-chart)
        - [Separate X and Y values](#separate-x-and-y-values)
        - [XY pairs](#xy-pairs)
        - [Y values only (X=index)](#y-values-only-xindex)
        - [Array data](#array-data)
    - [Line chart](#line-chart)
        - [Array data, separate X and Y values](#array-data-separate-x-and-y-values)
        - [Array data, XY pairs](#array-data-xy-pairs)
        - [Y values only (X=index)](#y-values-only-xindex-1)
    - [Scatter plot](#scatter-plot)
    - [Histogram](#histogram)
        - [Auto bin number](#auto-bin-number)
        - [Fixed bin number](#fixed-bin-number)
    - [CSV input](#csv-input)
- [Screenshots](#screenshots)
- [Licensing](#licensing)

<!-- /TOC -->

## Get it

```bash
go get -u github.com/sgreben/jp/cmd/jp
```

Or [download the binary](https://github.com/sgreben/jp/releases/latest) from the releases page. 

```bash
# Linux
curl -LO https://github.com/sgreben/jp/releases/download/1.1.7/jp_1.1.7_linux_x86_64.zip
unzip jp_1.1.7_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/jp/releases/download/1.1.7/jp_1.1.7_osx_x86_64.zip
unzip jp_1.1.7_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/jp/releases/download/1.1.7/jp_1.1.7_windows_x86_64.zip
unzip jp_1.1.7_windows_x86_64.zip
```

## Use it

`jp` reads JSON on stdin and prints plots to stdout.

```text
Usage of jp:
  -type value
    	Plot type. One of [line bar scatter hist] (default line)
  -x string
    	x values (JSONPath expression)
  -y string
    	y values (JSONPath expression)
  -xy string
    	x,y value pairs (JSONPath expression). Overrides -x and -y if given.
  -height int
    	Plot height (default 0 (auto))
  -width int
    	Plot width (default 0 (auto))
  -canvas value
    	Canvas type. One of [full quarter braille auto] (default auto)
  -bins uint
        Number of histogram bins (default 0 (auto))
  -input value
        Input type. One of [json csv] (default json)
```

## Examples

### Bar chart

#### Separate X and Y values

```bash
$ cat examples/tcp-time.json | jp -x ..Label -y ..Count -type bar

         69                                                                     
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████      21                                                       
    █████████████ █████████████       7             2             1             
    █████████████ █████████████ ▄▄▄▄▄▄▄▄▄▄▄▄▄ ▁▁▁▁▁▁▁▁▁▁▁▁▁ ▁▁▁▁▁▁▁▁▁▁▁▁▁       
                                                                                
     46.85267ms    48.38578ms    49.91889ms     51.452ms     52.98511ms         
```


#### XY pairs

```bash
$ cat examples/tcp-time.json | jp -xy "..[Label,Count]" -type bar

         69                                                                     
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████      21                                                       
    █████████████ █████████████       7             2             1             
    █████████████ █████████████ ▄▄▄▄▄▄▄▄▄▄▄▄▄ ▁▁▁▁▁▁▁▁▁▁▁▁▁ ▁▁▁▁▁▁▁▁▁▁▁▁▁       
                                                                                
     46.85267ms    48.38578ms    49.91889ms     51.452ms     52.98511ms                
```

#### Y values only (X=index)

```bash
$ cat examples/tcp-time.json | jp -y ..Count -type bar

         69                                                                     
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████                                                               
    █████████████      21                                                       
    █████████████ █████████████       7             2             1             
    █████████████ █████████████ ▄▄▄▄▄▄▄▄▄▄▄▄▄ ▁▁▁▁▁▁▁▁▁▁▁▁▁ ▁▁▁▁▁▁▁▁▁▁▁▁▁       
                                                                                
          0             1             2             3             4             
```

#### Array data

```bash
$ echo '[[-3, 5], [-2, 0], [-1, 0.1], [0, 1], [1, 2], [2, 3]]' | jp -xy '[*][0, 1]' -type bar

         5                                                                      
    ███████████                                                                 
    ███████████                                                                 
    ███████████                                                      3          
    ███████████                                                 ▄▄▄▄▄▄▄▄▄▄▄     
    ███████████                                          2      ███████████     
    ███████████                              1      ███████████ ███████████     
    ███████████                 0.1     ▄▄▄▄▄▄▄▄▄▄▄ ███████████ ███████████     
    ███████████      0      ▁▁▁▁▁▁▁▁▁▁▁ ███████████ ███████████ ███████████     
                                                                                
        -3          -2          -1           0           1           2          
```

### Line chart

#### Array data, separate X and Y values

```bash
$ jq -n '[range(200)/20 | [., sin]]' | jp -x '[*][0]' -y '[*][1]'
  1.059955│         ▄▄▄▖                                       ▗▄▄▄▖
          │       ▄▀▘  ▝▜▖                                   ▗▞▘   ▝▚
          │      ▟       ▝▄                                 ▗▀       ▜
          │     ▟         ▝▄                               ▗▀         ▜
          │    ▐           ▝▖                             ▗▞           ▚▖
          │   ▗▘            ▝▖                            ▞             ▚
          │  ▗▘              ▚                           ▞               ▌
          │  ▌                ▌                         ▗▘               ▝▖
          │ ▞                 ▝▖                        ▌                 ▚
          │▗▘                  ▚                       ▞                   ▌
          │▌                    ▌                     ▗▘                   ▝▖
          │                     ▝▖                   ▗▘                     ▐
          │                      ▐                   ▞                       ▚
          │                       ▚                 ▐                        ▝▖
          │                       ▝▖               ▗▘                         ▐
          │                        ▐               ▞                           ▚
          │                         ▀▖            ▐
          │                          ▚           ▄▘
          │                           ▙         ▗▘
          │                            ▚       ▄▘
          │                             ▚▄   ▗▞▘
          │                              ▝▀▀▀▘
 -1.059955└─────────────────────────────────────────────────────────────────────
          0                                                                 9.95
```

#### Array data, XY pairs

```bash
$ jq -n '[range(200)/20 | [., sin]]' | jp -xy '[*][0, 1]'
  1.059955│         ▄▄▄▖                                       ▗▄▄▄▖
          │       ▄▀▘  ▝▜▖                                   ▗▞▘   ▝▚
          │      ▟       ▝▄                                 ▗▀       ▜
          │     ▟         ▝▄                               ▗▀         ▜
          │    ▐           ▝▖                             ▗▞           ▚▖
          │   ▗▘            ▝▖                            ▞             ▚
          │  ▗▘              ▚                           ▞               ▌
          │  ▌                ▌                         ▗▘               ▝▖
          │ ▞                 ▝▖                        ▌                 ▚
          │▗▘                  ▚                       ▞                   ▌
          │▌                    ▌                     ▗▘                   ▝▖
          │                     ▝▖                   ▗▘                     ▐
          │                      ▐                   ▞                       ▚
          │                       ▚                 ▐                        ▝▖
          │                       ▝▖               ▗▘                         ▐
          │                        ▐               ▞                           ▚
          │                         ▀▖            ▐
          │                          ▚           ▄▘
          │                           ▙         ▗▘
          │                            ▚       ▄▘
          │                             ▚▄   ▗▞▘
          │                              ▝▀▀▀▘
 -1.059955└─────────────────────────────────────────────────────────────────────
          0                                                                 9.95
```

#### Y values only (X=index)

```bash
$ cat examples/tcp-time.json | jp -y ..Duration
 5.726165e+07│
             │
             │
             │
             │ ▗
             │ ▟
             │ █
             │▐▝▖
             │▐ ▌                                   ▌
             │▐ ▌                                   ▌
             │▌ ▌                                  ▐▚
             │▌ ▌                ▗       ▗         ▐▐    ▌
             │▘ ▌              ▖ ▐      ▞▀▖        ▐▐    ▌
             │  ▚   ▐▚  ▗▀▖   ▗▚ ▌▌    ▗▘ ▌ ▖▗▀▌   ▌▐    █
             │  ▐  ▛▌ ▚▖▞ ▚▐▖ ▞▐ ▌▌ ▗  ▐  ▐▟▐▞ ▚ ▗ ▌▝▖  ▐▐                      ▐
             │  ▐ ▐    ▝  ▝▌▝▀ ▝▟ ▚▗▜  ▞     ▘ ▐▖█▗▘ ▌  ▐▐    ▗  ▄▖     ▄▖      ▌▌
             │  ▝▚▐        ▘    ▘ ▐▘▝▖▄▌        ▝▝▟  ▀▀▚▟ ▌ ▖▞▘▌▐ ▚ ▗▄ ▐ ▚▄▖ ▄ ▄▘▌▞▄▄▀▚   ▄ ▄▗▞▖▞▄▄▚
             │   ▝▌                  ▛            ▌     ▝ ▙▞▝  ▝▘ ▝▚▘ ▀▘   ▝▀ ▀  ▐▘    ▚▞▀ ▀ ▘ ▚▘
             │
             │
             │
             │
 4.446018e+07└──────────────────────────────────────────────────────────────────────────────────────
             0                                                                                    99
```

### Scatter plot

```bash
$ cat examples/mvrnorm.json | jp -xy '..[x,y]' -type scatter

 3.535344│                                 ⠄             ⠄                     
         │                               ⠈⠂   ⠂       ⡀ ⠂                      
         │                          ⠐⡀⡀⡂   ⠁  ⢄  ⠁ ⠠                           
         │                            ⡀    ⠆     ⠈  ⠄⡀        ⠂                
         │           ⡀       ⠠  ⡀ ⡀ ⠄  ⡀⠐⠄⠁⠐ ⠠⢆⠠⠂⠂⠄⣀⢈  ⡀⠈ ⡀                    
         │                     ⡀⠂⠂⠄ ⡀⠂⢔⠠ ⢤⢀⠌⣡⠁⠦⠄⠐⡐⠂⣀⠅⠁⠈ ⠂ ⠈⠁⠁      ⡀     ⠄     
         │  ⡀         ⢀  ⠄     ⠈⠠ ⠡⠑⠈⠈⢢⡁⡄⢈⠂⢡⠈⡄⡀⠈⠰⢉⡠⠘⢄⢃⠉⢀⣄⠢⠠⠄ ⠠ ⡀⠁ ⡀ ⠂          
         │                   ⠈ ⡂⠈⡁⠈⠄⢂⡹⡐⡡⡆⡥⣙⡶⡼⠱⣅⣅⣼⢗⡱⢐⣈⠑⢁⠂ ⢐⢁⠭⠘⡀  ⠈              
         │                ⠁ ⢀⠄⢈⠈⡰⢀⡥⠋⣧⣓⣚⡛⢲⣽⣝⣭⢙⣟⢲⡽⣋⡠⣿⣜⣵⠙⡦⠗ ⣡⠁⠁⠁⠄⠠ ⠄⡂             
         │                  ⠄⠌⠌⠡⠉⡐⢯⣵⡏⢵⡞⠂⢰⣽⣷⢛⣯⡣⣷⢭⣞⣏⠤⣾⢡⡻⠢⢊⢠⡠⠸⢄⣃⡀⢁⠐ ⠐⡀ ⠂    ⠄     
         │              ⠨ ⡈⠂ ⢀⢑⠄⣜⡾⣴⢨⠶⣪⣧⢿⣷⣷⡱⣿⣞⣲⣮⣮⣯⢾⡷⡬⡷⣺⠤⢏⡼⣨⢌⡬⠠⢂⢠⠒⠱⠆⡈            
         │          ⠈  ⠃ ⠄⡐⠂⠐⢀⢈⣂⡈⣳⣷⣜⢺⣿⣹⣷⣼⣯⡿⣃⣽⣿⢾⣟⣾⢵⣻⠯⡼⡃⣼⣗⢲⠪⠇⣉⠺ ⢱⠠⠙⡀⢐⠌           
         │          ⠄   ⠈⠊⠐⠑⠨⠚⢁⡊⢾⡶⢩⢿⣏⣽⢞⣼⣇⣵⣿⣿⣽⣿⢽⣭⠺⣿⣽⣳⢚⣾⣻⣾⣜⠩⡒⣃⠈⢢⠕⢂⢰⡀  ⡔⢀⢀    ⡀   
         │            ⢀⠂⠁⠂⠇ ⠂⠊⢀⠐⡘⡍⡇⣚⢸⢟⣯⢿⣳⡪⣫⣵⣿⣯⣿⢿⣷⣻⣖⣗⣻⣚⢥⡷⣕⣏⠶⠊⠄⣠⠰⠂⡄⠂ ⠄⠁          
         │             ⠠⢂ ⠘⠐⣀ ⣀⢡⢐⠔⢫⠯⢕⠫⠿⣹⢶⣾⡻⣭⣽⢗⠿⣹⣛⣺⣿⠯⢲⡼⣵⢉⣭⢐⣟⡍⠄⠈⠥⠄   ⠁   ⠄       
         │             ⠈⠂  ⠊⡀⡈⠢⡌⡠⠖⢤⠥⡑⣯⣾⣴⣯⡿⣯⣝⣯⣿⠧⣽⣒⢾⣼⣻⣛⣗⡹⡽⢪⠯⠒⡨⠈ ⠈⡐⢄ ⠂⠘⠠        ⠂ 
         │              ⠰  ⡀⠃⠁⠠ ⠉⡈⡨⡱⢍⠌⠷⣯⠫⠬⡙⣴⣯⡣⡟⡮⠩⣫⠿⢞⢵⡰⠞⡂⠴ ⠕⢀⡂⠁ ⢀ ⠤   ⠈         
         │               ⢩⠂ ⠁⡄ ⢀⠲⢂⠑⢁⡘⠄⠵⣣⢑⢻⠨⡩⣌⠕⢮⣮⣋⢹⡁⣊⡃⠈⡕⡘⡠⠨⠄⡘⠨ ⠊⠁   ⠂           
         │          ⠐     ⢀ ⠈⠐⠔⠈  ⠁⢀⣀⡃⣊⢁⡘⠁⠛⠨ ⠒⡑⡀⠵⢙⠄⠡⠢⠃⠄⠋⠅ ⠥⠁⠠⢀ ⠄               
         │                  ⢀⢁ ⠆   ⠉⠁⠐ ⠄⠁⢑⡀⢀⠠⠑⢡⢊⠂⠑⠌⡅⠊⠄⠉⢈⡐  ⡀ ⠠   ⠂             
         │                      ⢀   ⠑ ⠂  ⡁ ⠌⢠⠈⠂⠄⠉⡃⠈⠄   ⠂⠠⠁ ⠄  ⢀ ⠠              
         │                             ⠐ ⠐ ⢁⠂⠂⠢⠠⠄⠔⠐       ⠁                    
         │                         ⠢  ⠁    ⠂⠐  ⠐       ⠐   ⠈                   
         │                           ⢀  ⠄⠈       ⠈                             
         │                                        ⠐                            
         │                                                                     
         │                                                                     
         │                                      ⠈                              
-4.271874└─────────────────────────────────────────────────────────────────────
          -4.08815                                                       3.79083
```

### Histogram

#### Auto bin number

```
$ cat examples/mvrnorm.json | jp -x ..x -type hist
                                    684                                     1  [-3.27033,-2.69856)
                                   █████▌                                   2  [-2.69856,-2.12679)
                                   █████▌ 624                               3  [-2.12679,-1.55502)
                                   ███████████                              4  [-1.55502,-0.983254)
                               557 ███████████                              5  [-0.983254,-0.411485)
                             ▐████████████████                              6  [-0.411485,0.160285)
                             ▐████████████████                              7  [0.160285,0.732054)
                             ▐████████████████                              8  [0.732054,1.30382)
                             ▐████████████████                              9  [1.30382,1.87559)
                             ▐████████████████ 404                          10 [1.87559,2.44736)
                             ▐█████████████████████▌                        11 [2.44736,3.01913)
                         314 ▐█████████████████████▌                        12 [3.01913,3.5909]
                        ▄▄▄▄▄▟█████████████████████▌
                        ███████████████████████████▌
                        ███████████████████████████▌
                        ███████████████████████████▌
                        ███████████████████████████▌ 176
                        █████████████████████████████████
                    98  █████████████████████████████████
                  ▐██████████████████████████████████████ 79
    1     4   41  ▐███████████████████████████████████████████  14    4
  ▁▁▁▁▁▁▁▁▁▁▁█████████████████████████████████████████████████▁▁▁▁▁▁▁▁▁▁▁

    0     1    2     3    4     5    6     7    8     9   10    11   12
```

#### Fixed bin number

```
$ cat examples/mvrnorm.json | jp -x ..x -type hist -bins 5
                                         1652
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████
                                    █████████████████       728
                                    ██████████████████████████████████
                          541       ██████████████████████████████████
                   ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄██████████████████████████████████
                   ███████████████████████████████████████████████████
                   ███████████████████████████████████████████████████
                   ███████████████████████████████████████████████████
                   ███████████████████████████████████████████████████
         22        ███████████████████████████████████████████████████       57
  ▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███████████████████████████████████████████████████▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄

  [-3.8421,-2.3555)[-2.3555,-0.8689)[-0.8689,0.6177)  [0.6177,2.1043)  [2.1043,3.5909]
```

### CSV input

```
$ cat examples/sin.csv | jp -input csv -xy '[*][0,1]'

  1.059955│       ▗▄▛▀▀▚▄▖                                    ▄▄▀▀▀▄▄
          │     ▗▞▘      ▝▚▖                                ▄▀      ▝▀▄
          │    ▟▘          ▝▄                             ▗▀          ▝▀▖
          │  ▗▛              ▚▖                          ▞▘             ▝▙
          │ ▄▘                ▀▖                        ▞                 ▚
          │▞▘                  ▝▌                     ▗▛                   ▚▖
          │                     ▝▚                   ▐▘                     ▝▄
          │                       ▜▖                ▟▘                       ▝▄
          │                        ▐▄             ▗▞                          ▝▚
          │                          ▚▖          ▄▀
          │                           ▀▙▖      ▄▛
          │                             ▀▀▄▄▄▞▀▘
 -1.059955└─────────────────────────────────────────────────────────────────────
          0                                                                 9.95
```

## Screenshots

In case you're on mobile, here's some PNGs of what `jp` output looks like:

- ![Scatter plot](docs/scatter_plot.png)
- ![Bar chart](docs/bar_chart.png)
- ![Line chart](docs/line_chart.png)

## Licensing

- Any original code is licensed under the [MIT License](./LICENSE).
- Included portions of [github.com/buger/goterm](https://github.com/buger/goterm) are licensed under the MIT License.
- Included portions of [github.com/kubernetes/client-go](https://github.com/kubernetes/client-go/tree/master/util/jsonpath) are licensed under the Apache License 2.0.
