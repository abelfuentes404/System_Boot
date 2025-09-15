import React from 'react';
import { useApi } from '../../hooks/useApi';
import { setupService } from '../../services/auth';

const CompleteSetup = ({ onBack }) => {
  const { loading, error, callApi } = useApi();

  const handleComplete = async () => {
    try {
      await callApi(() => setupService.complete());
      alert('Setup completado exitosamente! El sistema se recargará.');
      window.location.href = '/';
    } catch (err) {
      // Error manejado por useApi
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-6 text-center">Completar Setup</h2>
        
        {error && (
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
          </div>
        )}

        <div className="mb-6">
          <p className="text-gray-700 text-center">
            ¡Estás a punto de completar la configuración del sistema! 
            Se cifrarán y guardarán todas las configuraciones.
          </p>
        </div>

        <div className="flex gap-4">
          <button
            onClick={onBack}
            className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded"
          >
            Atrás
          </button>
          
          <button
            onClick={handleComplete}
            disabled={loading}
            className="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
          >
            {loading ? 'Completando...' : 'Completar Setup'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default CompleteSetup;