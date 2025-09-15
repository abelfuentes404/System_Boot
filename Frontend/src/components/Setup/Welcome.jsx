import React from 'react';

const Welcome = ({ onNext }) => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="text-center">
        <div className="mb-8">
          <h1 className="text-5xl font-bold mb-4">Bienvenido al Sistema</h1>
          <p className="text-xl opacity-80">Configuración inicial del sistema</p>
        </div>
        
        <div className="mb-8">
          <div className="inline-block relative">
            <div className="absolute -inset-4 bg-blue-500 rounded-lg blur opacity-30"></div>
            <div className="relative bg-gray-800 p-8 rounded-lg">
              <div className="text-6xl mb-4">🚀</div>
              <h2 className="text-2xl font-semibold mb-2">Sistema de Gestión</h2>
              <p className="opacity-70">Preparado para comenzar</p>
            </div>
          </div>
        </div>

        <button
          onClick={onNext}
          className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-8 rounded-lg text-lg transition duration-200 transform hover:scale-105"
        >
          Iniciar Configuración
        </button>
      </div>
    </div>
  );
};

export default Welcome;