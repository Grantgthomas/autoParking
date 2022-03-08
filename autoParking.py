from decimal import DivisionByZero
from multiprocessing.connection import deliver_challenge, wait
from operator import truediv
import sys
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time
import random
import sys
import getopt


def main(argv):
    # set path to driver
   
    #set browser to use
    driver = webdriver.Chrome("C:\Program Files (x86)\chromedriver.exe")
    #declare a string to represent path to form on website
    FName = "Karolina"
    LName = "Frankfurt"
    RNum = "999"
    VMake = ""
    VModel = ""
    VColor = ""
    VPlate = ""
    phoneNum = "8884475594"
    #alias email
    email = "u@s.co"
    values = [FName,LName,RNum,VMake,VModel,VColor,VPlate,phoneNum,email]
    try:
        opts,args = getopt.getopt(argv,"hi:o",["ma=","mo=","co=","pl=","e="])
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

    aptForm.send_keys("Luxia Swiss Ave (Encore Swiss Avenue)")
    firstNameForm.send_keys(values[0])
    lastNameForm.send_keys(values[1])
    unitForm.send_keys(values[2])
    makeForm.send_keys(VMake)
    modelForm.send_keys(VModel)
    colorForm.send_keys(VColor)
    plateForm.send_keys(VPlate)
    phoneForm.send_keys(values[7])
    emailForm.send_keys(email)
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
    #formSubmit.click()

if __name__ == "__main__":
    main(sys.argv[1:])

