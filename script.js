import Typed from typed.js
var typed = new Typed(".kontenerspan",{
    Strings: ["Comfort", "Stylish", "IMMACULATE!"],
    typeSpeed: 150,
    backSpeed: 150,
    looped: true
})


document.addEventListener("DOMContentLoaded", function () {
    // Get all product category elements
    var categories = document.querySelectorAll(".categories .card-title");

    // Get all product category containers
    var productCategories = document.querySelectorAll(".product .col");

    // Add click event listener to each category element
    categories.forEach(function (category) {
      category.addEventListener("click", function () {
        // Hide all product categories
        productCategories.forEach(function (productCategory) {
          productCategory.style.display = "none";
        });

        // Show the corresponding product category based on the clicked category
        var selectedCategory = category.textContent.toLowerCase();
        var selectedProductCategory = document.querySelector("." + selectedCategory);
        if (selectedProductCategory) {
          selectedProductCategory.style.display = "block";
        }
      });
    });
  });