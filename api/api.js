import axios from 'axios';

export default {
  data() {
    return {
      leftDrawerOpen: false,
      linksList: [],
      user: {}
    };
  },
  created() {
    this.fetchMockData();
  },
  methods: {
    async fetchMockData() {
      try {
        const response = await axios.get('../../api/server-mock.json');
        this.user = response.data.user;
      } catch (error) {
        console.error('Error fetching mock data:', error);
      }
    }
  }
};
