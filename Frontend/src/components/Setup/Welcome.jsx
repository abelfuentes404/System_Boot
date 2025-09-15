import React from 'react';

const Welcome = ({ onNext }) => {
  const handleClick = () => {
    console.log('Botón de iniciar configuración clickeado');
    if (onNext && typeof onNext === 'function') {
      onNext();
    } else {
      console.error('La función onNext no está definida o no es una función');
    }
  };

  return (
    <div className="min-h-screen stars-bg flex items-center justify-center">
      <div className="text-center text-white">
        <h1 className="text-5xl font-bold mb-6">Bienvenido al Sistema</h1>
        <p className="text-xl mb-8">Sistema de gestión integral</p>

        <div className="mb-8">
          <div className="inline-block relative">
            <div className="absolute -inset-4 bg-blue-500 rounded-lg blur opacity-30"></div>
            <div className="relative bg-gray-800 bg-opacity-80 p-8 rounded-lg backdrop-blur-sm">
              <div className="text-6xl mb-4">🚀</div>
              <h2 className="text-2xl font-semibold mb-2">Configuración Inicial</h2>
              <p className="opacity-70">Preparado para comenzar</p>
            </div>
          </div>
        </div>

        <button
          onClick={handleClick}
          className="relative z-50 bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-8 rounded-lg text-lg transition duration-200 transform hover:scale-105 focus:outline-none focus:ring-2 focus:ring-blue-300 active:scale-95 shadow-lg"
          style={{ cursor: 'pointer' }}
        >
          Iniciar Configuración
        </button>
      </div>
    </div>
  );
};

export default Welcome;