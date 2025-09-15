import React, { useState } from 'react';
import { useApi } from '../../hooks/useApi';
import { setupService } from '../../services/auth';

const DatabaseConfig = ({ onSuccess, onBack }) => {
  const [dbConfig, setDbConfig] = useState({
    host: 'localhost',
    port: 5432,
    user: 'postgres',
    password: '',
    dbname: 'mydatabase'
  });

  const { loading, error, callApi } = useApi();

  const handleTestConnection = async () => {
    try {
      await callApi(() => setupService.testConnection(dbConfig));
      alert('Conexión exitosa!');
    } catch (err) {
      // Error manejado por useApi
    }
  };

  const handleSave = async () => {
    try {
      await callApi(() => setupService.configure(dbConfig));
      onSuccess();
    } catch (err) {
      // Error manejado por useApi
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-6 text-center">Configuración de PostgreSQL</h2>
        
        {error && (
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
          </div>
        )}

        <form>
          {['host', 'port', 'user', 'password', 'dbname'].map((field) => (
            <div key={field} className="mb-4">
              <label className="block text-gray-700 text-sm font-bold mb-2 capitalize">
                {field}
              </label>
              <input
                type={field === 'password' ? 'password' : 'text'}
                value={dbConfig[field]}
                onChange={(e) => setDbConfig({ ...dbConfig, [field]: e.target.value })}
                className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
                required
              />
            </div>
          ))}

          <div className="flex gap-4">
            <button
              type="button"
              onClick={onBack}
              className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded"
            >
              Atrás
            </button>
            
            <button
              type="button"
              onClick={handleTestConnection}
              disabled={loading}
              className="bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
            >
              Probar Conexión
            </button>
            
            <button
              type="button"
              onClick={handleSave}
              disabled={loading}
              className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
            >
              Guardar
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default DatabaseConfig;