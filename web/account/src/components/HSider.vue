<template>
  <v-navigation-drawer floating color="rgb(0, 0, 0, 0)" app cliped>
    <v-flex xs12>
      <v-card height="100" color="rgb(0, 0, 0, 0)" flat></v-card>
    </v-flex>
    <v-list dense rounded>
      <v-list-item-group v-model="idx" color="primary">
        <template v-for="(item, i) in items">
          <template v-if="item.divider">
            <v-divider :key="item.title"></v-divider>
          </template>
          <template v-else>
            <v-list-item :href="item.href" :key="i">
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
        { title: "Home", icon: "account_circle", href: "/home" },
        {
          title: "Personal info",
          icon: "assignment_ind",
          href: "/person-info"
        },
        { title: "Data & personalization", icon: "toggle_on", href: "/" },
        { title: "Security", icon: "lock", href: "/" },
        { title: "People & sharing", icon: "people", href: "/" },
        { title: "Payments & subscriptions", icon: "payment", href: "/" },
        { divider: true },
        { title: "Help", icon: "help_outline", href: "/" },
        { title: "Send feedback", icon: "feedback", href: "/" }
      ]
    };
  }
};
</script>
