<template>
  <v-navigation-drawer floating permanent color="rgb(0, 0, 0, 0)">
    <v-list dense rounded>
      <v-list-item-group v-model="idx" color="primary">
        <template v-for="(item, i) in items">
          <template v-if="item.divider">
            <v-divider :key="item.title"></v-divider>
          </template>
          <template v-else>
            <v-list-item href="#" :key="i">
              <v-list-item-icon>
                <v-icon filled>{{ item.icon }}</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </template>
        </template>
      </v-list-item-group>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  data() {
    return {
      idx: 1,
      items: [
        { title: "Home", icon: "account_circle" },
        { title: "Personal info", icon: "fa-id-card" },
        { title: "Data & personalization", icon: "toggle_on" },
        { title: "Security", icon: "lock" },
        { title: "People & sharing", icon: "people" },
        { title: "Payments & subscriptions", icon: "payment" },
        { divider: true },
        { title: "Help", icon: "help_outline" },
        { title: "Send feedback", icon: "feedback" }
      ]
    };
  }
};
</script>
