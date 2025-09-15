import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Header from './components/Layout/Header';
import Home from './pages/Home';
import LoginPage from './pages/LoginPage';
import DashboardPage from './pages/DashboardPage';
import SetupPage from './pages/SetupPage';
import { setupService } from './services/auth';

function App() {
  const [setupStatus, setSetupStatus] = useState(null);
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState(null);

  // Verificar estado del setup y autenticación
  useEffect(() => {
    const checkSetupStatus = async () => {
      try {
        const status = await setupService.getStatus();
        setSetupStatus(status);
        
        // Verificar si hay usuario autenticado
        const token = localStorage.getItem('token');
        if (token) {
          // Aquí podrías verificar el token con el backend
          const userData = localStorage.getItem('user');
          if (userData) {
            setUser(JSON.parse(userData));
          }
        }
      } catch (error) {
        console.error('Error checking setup status:', error);
        setSetupStatus({ setupComplete: false, hasDbConfig: false });
      } finally {
        setLoading(false);
      }
    };

    checkSetupStatus();
  }, []);

  // Mostrar loading mientras se verifica el estado
  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-900">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  // Si el setup no está completo, redirigir a la página de setup
  if (setupStatus && !setupStatus.setupComplete) {
    return (
      <Router>
        <Routes>
          <Route path="/setup" element={<SetupPage />} />
          <Route path="*" element={<Navigate to="/setup" replace />} />
        </Routes>
      </Router>
    );
  }

  return (
    <Router>
      <div className="min-h-screen bg-gray-100 flex flex-col">
        <Header user={user} setUser={setUser} />
        <main className="flex-grow">
          <Routes>
            <Route path="/" element={<Home user={user} />} />
            <Route 
              path="/login" 
              element={
                user ? <Navigate to="/dashboard" replace /> : <LoginPage setUser={setUser} />
              } 
            />
            <Route 
              path="/dashboard" 
              element={
                user ? <DashboardPage /> : <Navigate to="/login" replace />
              } 
            />
            <Route path="/setup" element={<SetupPage />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;