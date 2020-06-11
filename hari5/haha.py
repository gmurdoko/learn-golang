lower = 20
upper = 30
print("Bilangan prima antara",lower,"and",upper,":") 
for num in range(lower,upper + 1): 
    # print(num)
    if num > 1: 
        for i in range(2,num): 
            # print(num,i)
            if (num % i) == 0: 
                # print(num,i)
                break 
            print(num,i)
