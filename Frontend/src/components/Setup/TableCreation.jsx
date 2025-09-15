import React from 'react';
import { useApi } from '../../hooks/useApi';
import { setupService } from '../../services/auth';

const TableCreation = ({ onSuccess, onBack }) => {
  const { loading, error, callApi } = useApi();

  const handleCreateTables = async () => {
    try {
      await callApi(() => setupService.createTables());
      onSuccess();
    } catch (err) {
      // Error manejado por useApi
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-6 text-center">Crear Tablas</h2>
        
        {error && (
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
          </div>
        )}

        <div className="mb-6">
          <h3 className="text-lg font-semibold mb-3">Tablas a crear:</h3>
          <ul className="list-disc list-inside text-gray-700">
            <li>smtp_config</li>
            <li>users</li>
            <li>reset_codes</li>
          </ul>
        </div>

        <div className="flex gap-4">
          <button
            onClick={onBack}
            className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded"
          >
            Atrás
          </button>
          
          <button
            onClick={handleCreateTables}
            disabled={loading}
            className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
          >
            {loading ? 'Creando...' : 'Crear Tablas'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default TableCreation;