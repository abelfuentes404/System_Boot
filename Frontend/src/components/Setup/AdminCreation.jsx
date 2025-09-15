import React, { useState } from 'react';
import { useApi } from '../../hooks/useApi';
import { setupService } from '../../services/auth';

const AdminCreation = ({ onSuccess, onBack }) => {
  const [adminData, setAdminData] = useState({
    email: '',
    password: ''
  });

  const { loading, error, callApi } = useApi();

  const handleCreateAdmin = async () => {
    try {
      await callApi(() => setupService.createAdmin(adminData));
      onSuccess();
    } catch (err) {
      // Error manejado por useApi
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-6 text-center">Crear Usuario Admin</h2>
        
        {error && (
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
          </div>
        )}

        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Email
          </label>
          <input
            type="email"
            value={adminData.email}
            onChange={(e) => setAdminData({ ...adminData, email: e.target.value })}
            className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
            required
          />
        </div>

        <div className="mb-6">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Contraseña
          </label>
          <input
            type="password"
            value={adminData.password}
            onChange={(e) => setAdminData({ ...adminData, password: e.target.value })}
            className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
            required
          />
        </div>

        <div className="flex gap-4">
          <button
            onClick={onBack}
            className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded"
          >
            Atrás
          </button>
          
          <button
            onClick={handleCreateAdmin}
            disabled={loading}
            className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
          >
            {loading ? 'Creando...' : 'Crear Admin'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default AdminCreation;