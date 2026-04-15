const express = require('express');
const axios = require('axios');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 3000;
const BACKEND_URL = process.env.BACKEND_URL || 'http://backend-service:5000';

app.use(express.static(path.join(__dirname, 'public')));

app.get('/api/data', async (req, res) => {
  try {
    const response = await axios.get(`${BACKEND_URL}/api/info`);
    res.json(response.data);
  } catch (error) {
    console.error('Error fetching data from backend:', error.message);
    res.status(500).json({ error: 'Failed to fetch data from backend' });
  }
});

app.get('/health', (req, res) => {
  res.json({ status: 'healthy', service: 'frontend' });
});

app.listen(PORT, () => {
  console.log(`Frontend server started on port ${PORT}`);
  console.log(`Backend URL: ${BACKEND_URL}`);
});