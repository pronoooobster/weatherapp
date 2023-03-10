# Weather app using the OpenWeatherMap API

## Installation and setup
1. Download the binary, coresponding to your OS, from the [available builds](https://github.com/pronoooobster/weatherapp/tree/main/build)
2. Ensure that the app could run as an executable
    - On Linux/MacOS, you need to make the binary executable by running the following command in the terminal
      ```
      chmod +x weatherapp
      ```
      or, when running on x86 MacOS, run with
      ```
      chmod +x weatherapp-x86
      ```
3. Run the binary 
    - On Windows, open the command prompt and run with 
      ```
      weatherapp.exe
      ```
    - On Linux/MacOS, open the terminal and run with 
      ```
      ./weatherapp
      ```
      or, when running on x86 MacOS, run with
      ```
      ./weatherapp-x86
      ```

## Usage
* The application will prompt you to enter a city name. Enter the city name and press enter. 
* The application will then fetch the weather data from the OpenWeatherMap API and display it in the terminal.
* Note that the application will not work if you do not have an internet connection. You will also get an error if the entered city name is not valid.
