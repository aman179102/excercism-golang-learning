package annalyn

// CanFastAttack returns true only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
    return !knightIsAwake
}

// CanSpy returns true if at least ONE of the 3 characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner returns true if prisoner is awake AND archer is asleep.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner returns true in 2 cases:
// 1. IF dog is present AND archer is sleeping → FREE
// 2. IF dog is NOT present → prisoner awake AND knight+archer both sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {

    // Case 1: Dog helps if archer is asleep
    if petDogIsPresent && !archerIsAwake {
        return true
    }

    // Case 2: No dog → prisoner must be awake, knight and archer must be asleep
    if !petDogIsPresent && prisonerIsAwake && !knightIsAwake && !archerIsAwake {
        return true
    }

    return false
}
