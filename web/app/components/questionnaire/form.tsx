import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCirclePlus, faBriefcase } from '@fortawesome/free-solid-svg-icons';
import { useState } from 'react';

export default function Form() {
  // useState
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState<Record<string, string>>({});
  const [experiences, setExperiences] = useState<Record<string, string>[]>([]);

  //Fonction
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e: any) => {
    e.preventDefault();
    setExperiences([...experiences, formData]);
    setFormData({});
    setShowForm(false);
  };

  const handleCancel = () => {
    setFormData({});
    setShowForm(false);
  };

  return (
    <>
      <section className='mb-8 border rounded-xl p-6'>
        <div className='flex items-center px-6 py-4'>
          <FontAwesomeIcon
            icon={faBriefcase}
            className='mr-4 text-indigo-600'
          />
          <h2 className='text-lg font-semibold'>Experience professionnelle</h2>
        </div>

        {experiences.map((exp, index) => (
          <div key={index} className='p-4 mb-4 border rounded bg-gray-50'>
            <h3 className='font-semibold mb-2'>Expérience:</h3>
            <p>{exp.test}</p>
          </div>
        ))}
        {showForm && (
          <form onSubmit={handleSubmit}>
            <label>Test</label>
            <input
              id='test'
              name='test'
              type='text'
              placeholder='test'
              value={formData.test || ''}
              onChange={handleChange}
            />

            <div>
              <button type='button' onClick={handleCancel}>
                Annuler
              </button>
              <button type='submit'>Valider</button>
            </div>
          </form>
        )}
        {!showForm && (
          <div className='flex justify-center py-4'>
            <button
              onClick={() => setShowForm(true)}
              type='button'
              className='inline-flex items-center px-6 py-2.5 text-sm font-semibold text-white bg-indigo-600 rounded-md hover:bg-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-600 transition-colors duration-200'
            >
              <FontAwesomeIcon icon={faCirclePlus} className='mr-2' />
              Ajouter une expérience professionnelle
            </button>
          </div>
        )}
      </section>
    </>
  );
}
