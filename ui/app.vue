<template>
  <div>
    <h1>Inventory Management System</h1>
    <UTable :rows="people" />
  </div>
</template>

<script setup>
const people = [
  {
    id: 1,
    name: "Laptop",
    description: "Dell",
    quantity: 1,
    price: 1.2,
  },
];
</script>
