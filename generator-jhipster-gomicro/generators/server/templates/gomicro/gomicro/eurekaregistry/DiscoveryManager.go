package eureka

func ManageDiscovery(configs RegistrationVariables)  {
		 manager := new (EurekaRegistrationManager)
		 manager.RegisterWithSerivceRegistry(configs)
		 manager.SendHeartBeat(configs)
}
