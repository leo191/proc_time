import re

def st_time(time_string):
    temp = re.findall(r'\d+', time_string) 
    res = list(map(int, temp)) 

    if "h" in time_string:
        hr_to_sec = res[0] * 60 *60
        mt_to_sec = res[1] * 60
        tot_sec = res[2] + mt_to_sec + hr_to_sec
    else:
        mt_to_sec = res[0] * 60
        tot_sec = res[1] + mt_to_sec

    print(str(tot_sec))

st_time("20m40")