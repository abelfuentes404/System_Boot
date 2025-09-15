import React from 'react';

const Dashboard = () => {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl font-bold mb-6">Dashboard</h1>
      
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-xl font-semibold mb-2">Usuarios</h2>
          <p className="text-3xl font-bold text-blue-500">15</p>
        </div>
        
        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-xl font-semibold mb-2">Productos</h2>
          <p className="text-3xl font-bold text-green-500">42</p>
        </div>
        
        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-xl font-semibold mb-2">Ventas</h2>
          <p className="text-3xl font-bold text-purple-500">1234</p>
        </div>
      </div>

      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold mb-4">Bienvenido al Sistema</h2>
        <p className="text-gray-700">
          El sistema se ha configurado correctamente y está listo para usar.
        </p>
      </div>
    </div>
  );
};

export default Dashboard;