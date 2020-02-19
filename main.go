package main

import (
	"time"

	"os"
	"os/signal"

	"syscall"

	"github.com/ophum/crd-test/pkg"
	clientset "github.com/ophum/crd-test/pkg/client/clientset/versioned"
	informers "github.com/ophum/crd-test/pkg/client/informers/externalversions"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	c := getConfig()
	kubeClient := getKubeClient(c)
	customClient := getCustomClient(c)
	customInformerFactory := informers.NewSharedInformerFactory(customClient, time.Second * 30)

	stopCh := SetupSignalHandler()
	go customInformerFactory.Start(stopCh)

	if err := ctrl.Run(R2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

getConfig() *rest.Config {
	c, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return c
}

func getKubeClient(c *rest.Config) kubernetes.Interface {
	cli, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func getCustomClient(c *rest.Config) clientset.Interface {
	cli, err := clientset.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
