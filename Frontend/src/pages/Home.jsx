import React from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';

const Home = () => {
  const { user } = useAuth();

  return (
    <div className="min-h-screen stars-bg flex items-center justify-center">
      <div className="text-center text-white">
        <h1 className="text-5xl font-bold mb-6">Bienvenido al Sistema</h1>
        <p className="text-xl mb-8">Sistema de gestión integral</p>
        
        {user ? (
          <Link
            to="/dashboard"
            className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-8 rounded-lg text-lg transition duration-200 transform hover:scale-105"
          >
            Ir al Dashboard
          </Link>
        ) : (
          <Link
            to="/login"
            className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-8 rounded-lg text-lg transition duration-200 transform hover:scale-105"
          >
            Iniciar Sesión
          </Link>
        )}
      </div>
    </div>
  );
};

export default Home;