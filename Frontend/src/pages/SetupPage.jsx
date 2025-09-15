import React, { useState } from 'react';
import Welcome from '../components/Setup/Welcome';
import DatabaseConfig from '../components/Setup/DatabaseConfig';
import TableCreation from '../components/Setup/TableCreation';
import AdminCreation from '../components/Setup/AdminCreation';
import CompleteSetup from '../components/Setup/CompleteSetup';

const SetupPage = () => {
  const [currentStep, setCurrentStep] = useState(0);

  const steps = [
    { component: Welcome, title: 'Bienvenida' },
    { component: DatabaseConfig, title: 'Configuración de BD' },
    { component: TableCreation, title: 'Creación de Tablas' },
    { component: AdminCreation, title: 'Usuario Admin' },
    { component: CompleteSetup, title: 'Completar Setup' },
  ];

  const CurrentComponent = steps[currentStep].component;

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