<template>
  <div v-if="items.length > 0">
    <h1>Inventory Management System</h1>
    <form @submit.prevent="onSubmitHandler(newItem)">
      <button>Add Item</button>
    </form>
    <UTable :rows="items" />
  </div>
</template>

<script setup>
const baseUrl = "http://localhost:8080";

const items = ref([]);
const newItem = ref({
  name: "Item 1",
  description: "Description",
  quantity: 1,
  price: 1.3,
});

onMounted(async () => {
  getAllItems();
});

async function onSubmitHandler(item) {
  await createItem(item);
  await getAllItems();
}

async function getAllItems(sortField = "name", sortOrder = "asc") {
  try {
    const { data } = await useFetch(
      `${baseUrl}/items?sortField=${sortField}&sortOrder=${sortOrder}`
    );
    items.value = data.value;
  } catch (error) {
    console.error(error);
  }
}

async function createItem(newItem) {
  try {
    const { data } = await useFetch(`${baseUrl}/items`, {
      method: "POST",
      body: newItem,
    });
  } catch (error) {
    console.error(error);
  }
}
</script>
