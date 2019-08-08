<template>
  <v-card class="mx-auto" flat outlined>
    <v-card height="150" flat>
      <v-flex xs12 row>
        <v-flex xs8 text-left ml-4>
          <v-card-title>
            <div>{{title}}</div>
          </v-card-title>
          <v-card-text>{{description}}</v-card-text>
        </v-flex>
        <v-flex xs3>
          <v-img class="mt-6 ml-4" :src="require('../../assets/img/'+ image)" width="96"></v-img>
        </v-flex>
      </v-flex>
    </v-card>

    <v-divider></v-divider>

    <v-card flat>
      <v-btn text width="100%" height="40" color="primary">{{linktext}}</v-btn>
    </v-card>
  </v-card>
</template>


<script>
export default {
  name: "HCard",
  props: {
    title: {
      type: String,
      default: "title",
      required: false
    },
    description: {
      type: String,
      default: "description",
      required: false
    },
    image: {
      type: String,
      default: "data_personalization.png",
      required: false
    },
    link: {
      type: String,
      default: "www.baidu.com",
      required: false
    },
    linktext: {
      type: String,
      default: "",
      required: true
    }
  }
  //   data: () => ({
  //     title: "Privacy & personalization"
  //   })
};
</script>
