import React, { useState, useEffect } from 'react';
import Welcome from '../components/Setup/Welcome';
import LocalLogin from '../components/Setup/LocalLogin';
import DatabaseConfig from '../components/Setup/DatabaseConfig';
import TableCreation from '../components/Setup/TableCreation';
import AdminCreation from '../components/Setup/AdminCreation';
import CompleteSetup from '../components/Setup/CompleteSetup';
import { setupService } from '../services/auth';

const SetupPage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  // Verificar si ya está autenticado
  useEffect(() => {
    const checkAuth = async () => {
      try {
        // Verificar si hay un token de setup válido
        const token = localStorage.getItem('setupToken');
        if (token) {
          setIsAuthenticated(true);
        }
      } catch (error) {
        console.error('Error checking auth:', error);
      }
    };

    checkAuth();
  }, []);

  const steps = [
    { component: Welcome, title: 'Bienvenida', requiresAuth: false },
    { component: LocalLogin, title: 'Login Local', requiresAuth: false },
    { component: DatabaseConfig, title: 'Configuración de BD', requiresAuth: true },
    { component: TableCreation, title: 'Creación de Tablas', requiresAuth: true },
    { component: AdminCreation, title: 'Usuario Admin', requiresAuth: true },
    { component: CompleteSetup, title: 'Completar Setup', requiresAuth: true },
  ];

  const handleNext = () => {
    if (currentStep < steps.length - 1) {
      setCurrentStep(currentStep + 1);
    }
  };

  const handleBack = () => {
    if (currentStep > 0) {
      setCurrentStep(currentStep - 1);
    }
  };

  const handleLoginSuccess = () => {
    setIsAuthenticated(true);
    handleNext(); // Avanzar al siguiente paso después del login exitoso
  };

  // Si el paso requiere autenticación y no está autenticado, redirigir al login
  if (steps[currentStep].requiresAuth && !isAuthenticated) {
    return <LocalLogin onSuccess={handleLoginSuccess} onBack={handleBack} />;
  }

  const CurrentComponent = steps[currentStep].component;

  return (
    <div>
      <CurrentComponent 
        onNext={handleNext} 
        onBack={handleBack}
        onSuccess={handleNext}
      />
    </div>
  );
};

export default SetupPage;