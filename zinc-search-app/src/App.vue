<template>
  <div id="app">
    <h1>Buscar en ZincSearch</h1>
    <label for="selector">Field:</label>
    <select @change="clean" v-model="selectedOption" id="selector">
      <option v-for="option in options" :key="option" :value="option">{{ option }}</option>
    </select>
    <input
      v-model="query"
      @input="search"
      type="text"
      placeholder="Escribe tu búsqueda..."
    />
    <div v-if="loading">Cargando...</div>
    <div v-if="error" class="error">
      <p>Error: {{ error }}</p>
    </div>
    <!-- <ul v-if="results.length">
      <li v-for="(result, index) in results" :key="index">
        <pre>{{ result }}</pre>
      </li>
    </ul> -->


    <h1 v-if="results.length">Tabla de respuestas</h1>
    
    <table v-if="results.length" border="1">
      <thead>
        <tr>
          <th>_Index</th>
          <th>_ID</th>
          <th>_Score</th>
          <th>Subject</th>
          <th>From</th>
          <th>To</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in results" :key="item._id">
          <td>{{ item._index }}</td>
          <td>{{ item._id }}</td>
          <td>{{ item._score }}</td>
          <td>{{ item._source.Subject }}</td>
          <td>{{ item._source.From }}</td>
          <td>{{ item._source.To }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  data() {
    return {
      query: "",
      results: [],
      loading: false,
      error: null,
      options: ["Body", "Date", "From", "Message-Id", "Subject", "To"],
      selectedOption: "Body" 
    };
  },
  methods: {
    async search() {
      if (!this.query) {
        this.results = [];
        return;
      }

      this.loading = true;
      this.error = null;

      try {
        const username = 'admin';  // Nombre de usuario
        const password = 'securepassword';  // Contraseña
        const response = await axios.get(`http://localhost:8080/search?${[this.selectedOption]}:${this.query}`,{
          auth: {
            username: username,
            password: password
          }
        });
        this.results = response.data.hits.hits;
        this.results = this.checkEmpty(this.results)
      } catch (err) {
        console.log(err)
        this.error = "No se pudo realizar la búsqueda.";
      } finally {
        this.loading = false;
      }
    },
    clean(){
      this.results = [];
    },
    checkEmpty(results){
      return results.map(item=>{
        const fields = Object.values(item._source);
        const filled = fields.filter(f=>f==="")
        if(filled.length === 7 || filled.length === 6) return null
        return item
      }).filter(p=>p) ;
    }
  },
};
</script>

<style>
#app {
  text-align: center;
  padding: 20px;
}

input {
  padding: 10px;
  font-size: 16px;
}

select{
  padding: 12px;
  margin-right: 15px;
}

.error {
  color: red;
}

ul {
  list-style-type: none;
  padding: 0;
}

pre {
  text-align: left;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
