# Details on this project
This Go program is designed to manage product information by allowing the user to add, view, edit, and remove products through a simple command-line interface. 
It starts by declaring three slices (names, prices, and stocks) to store product names, prices, and stock levels, respectively. 
The program runs in an infinite loop, displaying a menu with six options: adding a product, showing a product by name, removing a product, editing a product, listing all products, and exiting the program. 
Although only the "add product" functionality is partially implemented, the program uses fmt.Scan to gather user input and store it in the appropriate slices. 
The other menu options are placeholders with no functionality yet, and there are also repeated print statements labeling input as "Enter the name of the new product," which should be revised for price and stock. The program breaks out of the loop when the user chooses to exit by selecting option 6.
