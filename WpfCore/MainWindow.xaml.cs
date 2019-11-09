using System.Windows;

namespace WpfCore
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();

            var parentProcess = ParentProcessUtilities.GetParentProcess();
            parentProcess.Kill();
        }
    }
}
