import api from './api';

export const authService = {
  login: async (credentials) => {
    const response = await api.post('/auth/login', credentials);
    return response.data;
  },

  logout: async () => {
    const response = await api.post('/auth/logout');
    return response.data;
  },

  verifyToken: async () => {
    const response = await api.get('/auth/verify');
    return response.data;
  },
};

export const setupService = {
  getStatus: async () => {
    const response = await api.get('/setup/status');
    return response.data;
  },

  login: async (credentials) => {
    const response = await api.post('/setup/login', credentials);
    return response.data;
  },

  testConnection: async (dbConfig) => {
    const response = await api.post('/setup/test-connection', dbConfig);
    return response.data;
  },

  configure: async (dbConfig) => {
    const response = await api.post('/setup/configure', dbConfig);
    return response.data;
  },

  createTables: async () => {
    const response = await api.post('/setup/create-tables');
    return response.data;
  },

  createAdmin: async (adminData) => {
    const response = await api.post('/setup/create-admin', adminData);
    return response.data;
  },

  complete: async () => {
    const response = await api.post('/setup/complete');
    return response.data;
  },
};