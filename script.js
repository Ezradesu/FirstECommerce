import Typed from typed.js
var typed = new Typed(".kontenerspan",{
    Strings: ["Comfort", "Stylish", "IMMACULATE!"],
    typeSpeed: 150,
    backSpeed: 150,
    looped: true
})


document.addEventListener("DOMContentLoaded", function () {
    // kategori element
    var categories = document.querySelectorAll(".categories .card-title");

    // element kontener
    var productCategories = document.querySelectorAll(".product .col");

    
    categories.forEach(function (category) {
      category.addEventListener("click", function () {
        // Hide all product categories
        productCategories.forEach(function (productCategory) {
          productCategory.style.display = "none";
        });

       
        var selectedCategory = category.textContent.toLowerCase();
        var selectedProductCategory = document.querySelector("." + selectedCategory);
        if (selectedProductCategory) {
          selectedProductCategory.style.display = "block";
        }
      });
    });
  });

  document.addEventListener("DOMContentLoaded", function () {
    
    var products = document.querySelectorAll(".product .card");
    var cartItemsContainer = document.getElementById("cartItems");

    products.forEach(function (product) {
        product.addEventListener("click", function () {
            var clonedProduct = product.cloneNode(true);
            cartItemsContainer.appendChild(clonedProduct);
        });
    });

    // buat masukin email
    var checkoutForm = document.getElementById("checkoutForm");
    checkoutForm.addEventListener("submit", function (event) {
        event.preventDefault();

        // Get email dari user
        var userEmail = document.getElementById("email").value;

        // masukin value produk ke cart
        var selectedProductNames = [];
        var cartItems = cartItemsContainer.querySelectorAll(".card-title");
        cartItems.forEach(function (item) {
            selectedProductNames.push(item.textContent.trim());
        });

        // nampilin yang dipilih
        console.log("Email: " + userEmail);
        console.log("Selected Products: " + selectedProductNames.join(", "));
    });
});