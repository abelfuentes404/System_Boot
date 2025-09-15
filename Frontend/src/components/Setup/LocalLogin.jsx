import React, { useState } from 'react';
import { useApi } from '../../hooks/useApi';
import { setupService } from '../../services/auth';

const LocalLogin = ({ onSuccess, onBack }) => {
  const [credentials, setCredentials] = useState({
    username: 'admin',
    password: 'admin123'
  });

  const { loading, error, callApi } = useApi();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await callApi(() => 
        setupService.login(credentials)
      );
      
      // Guardar token temporal para el setup
      localStorage.setItem('setupToken', response.token);
      onSuccess();
    } catch (err) {
      // Error ya manejado por useApi
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="bg-gray-800 p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-6 text-center">Login de Configuración</h2>
        
        {error && (
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
          </div>
        )}

        <form onSubmit={handleLogin}>
          <div className="mb-4">
            <label className="block text-gray-300 text-sm font-bold mb-2">
              Usuario
            </label>
            <input
              type="text"
              value={credentials.username}
              onChange={(e) => setCredentials({ ...credentials, username: e.target.value })}
              className="w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded focus:outline-none focus:border-blue-500 text-white"
              required
            />
          </div>

          <div className="mb-6">
            <label className="block text-gray-300 text-sm font-bold mb-2">
              Contraseña
            </label>
            <input
              type="password"
              value={credentials.password}
              onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
              className="w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded focus:outline-none focus:border-blue-500 text-white"
              required
            />
          </div>

          <div className="flex gap-4">
            <button
              type="button"
              onClick={onBack}
              className="bg-gray-600 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded"
            >
              Atrás
            </button>
            
            <button
              type="submit"
              disabled={loading}
              className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded disabled:opacity-50 flex-1"
            >
              {loading ? 'Iniciando sesión...' : 'Iniciar Sesión'}
            </button>
          </div>
        </form>

        <div className="mt-6 p-4 bg-gray-700 rounded">
          <p className="text-sm text-gray-300">
            <strong>Credenciales predeterminadas:</strong><br />
            Usuario: admin<br />
            Contraseña: admin123
          </p>
        </div>
      </div>
    </div>
  );
};

export default LocalLogin;