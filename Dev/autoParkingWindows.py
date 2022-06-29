from decimal import DivisionByZero
from multiprocessing.connection import deliver_challenge
from operator import truediv
import sys
from requests import options
from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.keys import Keys
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.firefox.service import Service as FirefoxService
from selenium.webdriver.chrome.service import Service as ChromeService
from selenium.webdriver.firefox.firefox_binary import FirefoxBinary
from selenium.webdriver.chrome.options import Options
from webdriver_manager.chrome import ChromeDriverManager
from webdriver_manager.firefox import GeckoDriverManager

import time
import random
import sys
import getopt


def main(argv):
    # set path to driver
   

    #Web driver for Docker container
    #driver = webdriver.Remote("http://127.0.0.1:4444/wd/hub", DesiredCapabilities.CHROME,options=set_chrome_options)
    
    #set the driver manager
    #service = ChromeService(executable_path=ChromeDriverManager().install())
    binary = FirefoxBinary('usr/bin/firefox')
    service = FirefoxService(executable_path='/usr/bin/firefox')

    #set browser to use
    #service = Service(executable_path=r"C:\\Users\\gamad\Documents\\Projects\\AutoParking")
    #driver = webdriver.Chrome(service=service,options=set_chrome_options)
    
    #Firefox driver alternative
    
    driver = webdriver.Firefox(service=service)


    #declare a string to represent path to form on website
    FName = "Karolina"
    LName = "Frankfurt"
    RNum = "999"
    VMake = ""
    VModel = ""
    VColor = ""
    VPlate = ""
    phoneNum = "8884475594"
    apartment = ""
    #alias email
    email = "u@s.co"
    values = [FName,LName,RNum,VMake,VModel,VColor,VPlate,phoneNum,email]
    try:
        opts,args = getopt.getopt(argv,"hi:o",["ma=","mo=","co=","pl=","e=","apt="])
    except getopt.GetoptError:
        print('options failed to load')
        sys.exit(2)      
#    opts,args = getopt.getopt(args,"hi:o",["make=","model=","color=","plate="])

    for opt, arg in opts:
        if opt in '--ma':
            VMake = arg
        elif opt in '--mo':
            VModel = arg
        elif opt in '--co':
            VColor = arg
        elif opt in '--pl':
            VPlate = arg
        elif opt in '--e':
            email = arg
        elif opt in '--apt':
            apartment =arg

        #Load input forms
    driver.get("https://app.parkingbadge.com/#/guest")

    try:
        aptForm = WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.XPATH, "/html/body/div/div[2]/form/div[1]/div/div/div/input")))
    except:
        driver.quit

    aptForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[2]/form/div[1]/div/div/div/input")
    firstNameForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,2,1))
    lastNameForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,2,2))
    unitForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,2,3))
    makeForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,3,1))
    modelForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,3,2))
    colorForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,3,3))
    plateForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,3,4))
    phoneForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,4,1))
    emailForm = driver.find_element(by=By.XPATH, value="/html/body/div/div[{}]/form/div[{}]/div/div[{}]/div/div/input".format(2,4,2))
    try:
        aptForm.send_keys(apartment)
        firstNameForm.send_keys(values[0])
        lastNameForm.send_keys(values[1])
        unitForm.send_keys(values[2])
        makeForm.send_keys(VMake)
        modelForm.send_keys(VModel)
        colorForm.send_keys(VColor)
        plateForm.send_keys(VPlate)
        phoneForm.send_keys(values[7])
        emailForm.send_keys(email)
        time.sleep(random.uniform(2.000000,4.988888))
        try:
            aptSelect = WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.XPATH, "/html/body/div/div[2]/form/div[1]/div/div/div/ul"))
            )
        except:
            driver.quit
        aptSelect.click()

        acceptBox = driver.find_element(by=By.CLASS_NAME, value="checkmark")
        if (acceptBox.is_selected):
            print("")
        else:
            acceptBox.click()


        time.sleep(random.uniform(2.000000,4.988888))
        formSubmit = driver.find_element(by=By.XPATH, value="/html/body/div/div[2]/form/div[7]")
        formSubmit.click()
        time.sleep(random.uniform(2.000000,4.988888))
        time.sleep(random.uniform(2.000000,4.988888))
        formSubmit.click()
        time.sleep(10) 
        driver.close()
        driver.quit()
        driver.close()
    except:
        pass
    else:
        driver.close()
        driver.quit()
    finally:
        driver.close()
        driver.quit()
        driver.close()
        

    
def set_chrome_options():
    chrome_options = Options()
    chrome_options.add_argument("--headless")
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_prefs = {}
    chrome_options.experimental_options["prefs"] = chrome_prefs
    chrome_prefs["profile.default_content_settings"] = {"images": 2}
    return chrome_options

def set_firefox_options():
    firefox_options = Options()
    firefox_options.add_argument("--headless")
    firefox_options.add_argument("--no-sandbox")
    firefox_options.add_argument("--disable-dev-shm-usage")    
    return firefox_options


if __name__ == "__main__":
    main(sys.argv[1:])

